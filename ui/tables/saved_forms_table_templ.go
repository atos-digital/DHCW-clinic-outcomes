// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package tables

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"github.com/atos-digital/DHCW-clinic-outcomes/internal/store/db"
)

func SavedFormsTable(saves []db.Save) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
<<<<<<< HEAD
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-get=\"/outcomes\" hx-target=\"body\" class=\"dhcw-border w-20 p-1 text-white bg-sky-700\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := `New`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button><table class=\"table-fixed border border-gray-300 w-1/2\"><tr class=\"bg-gray-300\"><td class=\"p-2 font-medium w-10\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var3 := `ID`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"p-2 font-medium \">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := `Date Created`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"p-2 font-medium w-32\"></td></tr>")
=======
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-get=\"/outcomes\" hx-target=\"body\" class=\"dhcw-border w-20 p-1 text-white bg-sky-700\">New</button><table class=\"table-fixed border border-gray-300 w-2/3\"><tr class=\"bg-gray-300\"><td class=\"p-2 font-medium w-10\">ID</td><td class=\"p-2 font-medium \">Date Created</td><td class=\"p-2 font-medium\">Speciality</td><td class=\"p-2 font-medium w-1/3\">Senior Responsible Clinician</td><td class=\"p-2 font-medium w-32\"></td></tr>")
>>>>>>> origin/master
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, save := range saves {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr><td class=\"p-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
<<<<<<< HEAD
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(save.ID)
=======
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(save.ID)
>>>>>>> origin/master
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/tables/saved_forms_table.templ`, Line: 17, Col: 29}
			}
<<<<<<< HEAD
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
=======
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
>>>>>>> origin/master
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"p-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
<<<<<<< HEAD
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(save.DateCreated.Format("02-Jan-06 15:04 MST"))
=======
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(save.DateCreated.Format("02-Jan-06 15:04 MST"))
>>>>>>> origin/master
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/tables/saved_forms_table.templ`, Line: 18, Col: 68}
			}
<<<<<<< HEAD
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
=======
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"p-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(save.Data.EventSpecialty)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/tables/saved_forms_table.templ`, Line: 21, Col: 46}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"p-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(save.Data.EventClinician)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/tables/saved_forms_table.templ`, Line: 22, Col: 46}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
>>>>>>> origin/master
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"p-2\"><button class=\"bg-sky-700 px-3 py-1 text-white rounded-3xl shadow-sm hover:bg-sky-600\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("/hx/load-save/%s", save.ID)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
<<<<<<< HEAD
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"/outcomes\" hx-target=\"body\" hx-swap=\"show:window:top\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var7 := `Load Draft`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var7)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></td></tr>")
=======
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"/outcomes\" hx-target=\"body\" hx-swap=\"show:window:top\">Load Draft</button></td></tr>")
>>>>>>> origin/master
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</table>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
