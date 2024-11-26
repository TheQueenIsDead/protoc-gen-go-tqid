package pkg

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"google.golang.org/protobuf/compiler/protogen"
	"log"
	"strings"
)

func Generate(plugin *protogen.Plugin, opts Options) (err error) {

	err = generateBoilerplate(plugin, opts)
	if err != nil {
		log.Println("Error generating boilerplate:", err)
		return
	}

	err = generateMessages(plugin, opts)
	if err != nil {
		log.Println("Error generating messages:", err)
		return
	}

	err = generateServices(plugin, opts)
	if err != nil {
		log.Println("Error generating services:", err)
		return
	}

	return
}

func generateBoilerplate(plugin *protogen.Plugin, opts Options) (err error) {

	// Read main.go and interpolate the servicename where required.
	buf, err := ReadFsFile(Main)
	repl := strings.Replace(string(buf), `"template"`, fmt.Sprintf(`"%s"`, opts.ServiceName), -1)
	buf = []byte(repl)

	// WIP Attempt to parse the given file into it's AST
	fest := token.NewFileSet()
	f, err := parser.ParseFile(fest, "", buf, parser.ParseComments)
	if err != nil {
		return err
	}
	log.Println(f)
	ast.Inspect(f, func(n ast.Node) bool {
		funcDecl, ok := n.(*ast.FuncDecl)
		if ok {
			if funcDecl.Name.Name == "main" {
				log.Println("FOUND MAIN:")
				log.Println(funcDecl.Body.List[0])
			}
		}
		return true
	})
	// END: WIP Attempt to parse the given file into it's AST

	// Write a new generated file to the plugin for rendering later
	file := plugin.NewGeneratedFile("main.go", ".")
	_, err = file.Write(buf)
	if err != nil {
		log.Println("Error writing template file:", err)
		return
	}

	return
}

func generateMessages(plugin *protogen.Plugin, opts Options) (err error) {

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
		// TODO: Try newGeneratedFile.P() to print instead?
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

func generateServices(plugin *protogen.Plugin, opts Options) (err error) {
	return
}
