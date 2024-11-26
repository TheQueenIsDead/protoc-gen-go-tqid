package pkg

import (
	"bytes"
	"io"
	"log"
	"os"

	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

func Generate(plugin *protogen.Plugin) (err error) {

	err = generateBoilerplate(plugin)
	if err != nil {
		log.Println("Error generating boilerplate:", err)
		return
	}

	err = generateMessages(plugin)
	if err != nil {
		log.Println("Error generating messages:", err)
		return
	}

	err = generateServices(plugin)
	if err != nil {
		log.Println("Error generating services:", err)
		return
	}

	return
}

func generateBoilerplate(plugin *protogen.Plugin) (err error) {

	f, err := os.Open("./pkg/template/main.go")
	buf, err := io.ReadAll(f)
	if err != nil {
		log.Println("Error reading template file:", err)
		return
	}

	file := plugin.NewGeneratedFile("main.go", ".")
	_, err = file.Write(buf)
	if err != nil {
		log.Println("Error writing template file:", err)
		return
	}

	return
}

func generateMessages(plugin *protogen.Plugin) (err error) {

	// Protoc passes a slice of File structs for us to process
	for _, file := range plugin.Files {

		// Time to generate code...!

		// 1. Initialise a buffer to hold the generated code
		var buf bytes.Buffer

		// 2. Write the package name
		//log.Println(newGeneratedFile)
		pkg := fmt.Sprintf("package %s", file.GoPackageName)
		buf.Write([]byte(pkg))

		// 3. For each message add our Foo() method
		for _, msg := range file.Proto.MessageType {
			buf.Write([]byte(fmt.Sprintf(`
            func (x %s) Foo() string {
               return "bar"
            }`, *msg.Name)))
		}

		// 4. Specify the output filename
		filename := file.GeneratedFilenamePrefix + ".tqid.pb.go"
		newGeneratedFile := plugin.NewGeneratedFile(filename, ".")

		// 5. Pass the data from our buffer to the plugin newGeneratedFile struct
		write, err := newGeneratedFile.Write(buf.Bytes())
		if err != nil {
			log.Println("Error writing generated file:", err)
			return err
		} else if write != len(buf.Bytes()) {
			log.Println("Error writing generated file: did not write all of them")
			return ErrBadWrite
		}
	}

	return
}

func generateServices(plugin *protogen.Plugin) (err error) {
	return
}
