package templates

// dtoTmpl : template of DTO
var dtoTmpl = `package dto
import ( 
	{{ range $key, $value := .Import}}"{{$key}}"
	{{ end }}
	"{{.Module}}/dal"
)  

// {{DTO .StructName}} : data transfer object {{ if .IsTable}} ({{.TableName}}) table.{{ else if .IsView}} ({{.ViewName}}) view.{{ end }}
type {{DTO .StructName}} struct {
	{{range .Fields}}{{.DTOfmt}}
	{{end}}
}

// {{DTO .StructName}}ToDAL : convert {{DTO .StructName}} to {{DAL .StructName}}
func (a *{{DTO .StructName}}) {{DTO .StructName}}ToDAL() (*dal.{{DAL .StructName}}, error) { 
	{{toLower .StructName}} := &dal.{{DAL .StructName}}{
		{{range .Fields}}{{.GoName}}:a.{{.GoName}},
		{{end}} 
	}
	return {{toLower .StructName}}, nil
}

// {{DAL .StructName}}ToDTO : convert {{DAL .StructName}} to {{DTO .StructName}}
func {{DAL .StructName}}ToDTO(a *dal.{{DAL .StructName}}) (*{{DTO .StructName}}, error) { 
	{{toLower .StructName}} := &{{DTO .StructName}}{
		{{range .Fields}}{{.GoName}}:a.{{.GoName}},
		{{end}} 
	}
	return {{toLower .StructName}}, nil
}

// {{DAL .StructName}}ToDTOArr : convert Array of {{DAL .StructName}} to Array of {{DTO .StructName}}
func {{DAL .StructName}}ToDTOArr({{pluralizeLower .StructName}} []*dal.{{DAL .StructName}}) ([]*{{DTO .StructName}}, error) {
	var err error
	res := make([]*{{DTO .StructName}}, len({{pluralizeLower .StructName}}))
	for i, {{toLower .StructName}} := range {{pluralizeLower .StructName}} {
		res[i], err = {{DAL .StructName}}ToDTO({{toLower .StructName}})
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


`
