package cmd

import (
	"log"
	"os"
	"path"
	"text/template"

	"github.com/mitoteam/mt-tgadmin/app"
	"github.com/mitoteam/mt-tgadmin/mttools"

	"github.com/spf13/cobra"
)

const systemdServiceDirPath = "/usr/lib/systemd"

func init() {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Creates system service to run " + app.Global.AppName,

		RunE: func(cmd *cobra.Command, args []string) error {
			//log.Printf("Autostart %t\n", app.Global.ServiceAutostart)

			if mttools.IsDirExists(systemdServiceDirPath) {
				systemdServiceInstall()
			} else {
				log.Fatalf(
					"Directory %s does not exists. Only systemd based services supported for now.\n",
					systemdServiceDirPath,
				)
			}

			return nil
		},
	}

	cmd.PersistentFlags().BoolVar(
		&app.Global.ServiceAutostart,
		"autostart",
		false,
		"Set service to be auto started after boot. Please note this does not auto starts service after installation.",
	)

	rootCmd.AddCommand(cmd)
}

type serviceUnitData struct {
	Name       string
	User       string
	Group      string
	Executable string
	WorkingDir string
}

func systemdServiceInstall() {
	filename := path.Join(systemdServiceDirPath, app.Global.Settings.ServiceName+".service")

	if mttools.IsFileExists(filename) {
		log.Fatalf("File %s already exists. Use 'uninstall' command or remove file manually.\n", filename)
	}

	t := template.New("service")
	if _, err := t.Parse(app.ServiceUnitFileTemplate); err != nil {
		log.Fatalln(err)
	}

	executable, err := os.Executable()

	if err != nil {
		log.Fatalln(err)
	}

	workingDir, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	data := &serviceUnitData{
		Name:       app.Global.Settings.ServiceName,
		User:       app.Global.Settings.ServiceUser,
		Group:      app.Global.Settings.ServiceGroup,
		Executable: executable,
		WorkingDir: workingDir,
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}

	if err := t.Execute(file, data); err != nil {
		log.Fatalln(err)
	}

	file.Close()

	if err := file.Chmod(0644); err != nil {
		log.Fatal(err)
	}
}
