package mailer

import (
	"bytes"
	"errors"
	"html/template"
	"net/mail"
	"os"
	"path/filepath"

	"github.com/farkmi/pocketbase-templ-starter/internal/server/config"
	"github.com/farkmi/pocketbase-templ-starter/internal/util"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

type EmailTemplate string

func (t EmailTemplate) String() string {
	return string(t)
}

var (
	WebTemplatesEmailBaseDirAbs               = "/web/templates/email"
	ErrEmailTemplateNotFound                  = errors.New("email template not found")
	ErrEmailAssetURL                          = errors.New("email asset url error")
	emailTemplateHello          EmailTemplate = "list_download" // /app/web/templates/email/hello/**.
	mailerInstance              *Mailer
	emailLogoFilename           = "logo.png"
)

type Mailer struct {
	app         *pocketbase.PocketBase
	TemplateDir string
	Templates   map[string]*template.Template
}

func Init(app *pocketbase.PocketBase) error {
	templateDir := config.GetServerConfig().EmailTemplatesDir

	m := &Mailer{
		app:         app,
		TemplateDir: templateDir,
		Templates:   map[string]*template.Template{},
	}

	if err := m.ParseTemplates(); err != nil {
		app.Logger().Error("Failed to parse email templates", "error", err)
		return err
	}

	mailerInstance = m

	return nil
}

func GetMailer() *Mailer {
	return mailerInstance
}

func (m *Mailer) ParseTemplates() error {
	log := m.app.Logger()

	files, err := os.ReadDir(m.TemplateDir)
	if err != nil {
		log.Error("Failed to read email templates directory while parsing templates",
			"dir", m.TemplateDir,
			"error", err,
		)
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		t, err := template.ParseGlob(filepath.Join(m.TemplateDir, file.Name(), "**"))
		if err != nil {
			log.Error("Failed to parse email template files as glob",
				"template", file.Name(),
				"error", err,
			)
			return err
		}

		m.Templates[file.Name()] = t
	}

	return nil
}

func (m *Mailer) SendHelloEmail(to string) error {
	log := m.app.Logger().With("template", emailTemplateHello).With("recipient", to)

	t, ok := m.Templates[string(emailTemplateHello)]
	if !ok {
		log.Error("Email template not found", "error", ErrEmailTemplateNotFound)
		return ErrEmailTemplateNotFound
	}

	data := map[string]interface{}{
		"logoUrl": util.PublicAssetsLink(emailLogoFilename),
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		log.Error("Failed to execute list download email template", "error", err)
		return err
	}

	message := &mailer.Message{
		From: mail.Address{
			Address: m.app.Settings().Meta.SenderAddress,
			Name:    m.app.Settings().Meta.SenderName,
		},
		To:      []mail.Address{{Address: to}},
		Subject: "Hello World!",
		HTML:    buf.String(),
	}

	return m.app.NewMailClient().Send(message)
}
