// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.898
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/crispyfishtech/crispyfish-demo/views/components"

func Index(title string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Header(title).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<body class=\"bg-gradient-to-br from-blue-500 via-cyan-400 to-orange-400 min-h-screen flex flex-col items-center justify-center\"><div class=\"w-full mx-auto p-8 h-screen flex items-center\"><div class=\"flex flex-col gap-8 md:flex-row w-full h-full\"><!-- Left Column --><div class=\"md:w-2/6 flex flex-col items-center bg-gray-900 rounded-xl shadow-lg p-8 h-full\"><div class=\"mb-4 text-gray-400 text-sm\">Inspired by <a href=\"https://github.com/oskapt/rancher-demo\" target=\"_blank\" class=\"underline hover:text-cyan-400\">oskapt/rancher-demo</a></div><img class=\"w-40 h-40 mb-6\" src=\"/static/img/cft-light.png\" alt=\"Crispyfish Logo\"><div class=\"mb-6 flex flex-col items-center\"><div class=\"text-5xl md:text-6xl font-extrabold text-cyan-400\" id=\"container-count\"></div><div class=\"text-lg md:text-xl font-bold text-gray-300 tracking-wide uppercase\" id=\"container-count-label\"></div></div><div class=\"mb-4 flex flex-col items-center container-backend hidden\" id=\"container-backend\"><div class=\"text-2xl md:text-3xl font-extrabold text-white text-center\" id=\"current-container\"></div><div class=\"text-lg md:text-xl font-bold text-gray-300 tracking-wide uppercase\">current backend</div></div><div class=\"mb-4 flex flex-row gap-6 items-center\"><div class=\"flex flex-col items-center\"><div class=\"text-3xl md:text-4xl font-extrabold text-cyan-400\" id=\"requests-count\">0</div><div class=\"text-lg md:text-xl font-bold text-gray-300 tracking-wide uppercase\">requests</div></div><div class=\"flex flex-col items-center\"><div class=\"text-3xl md:text-4xl font-extrabold text-red-500\" id=\"requests-error-count\">0</div><div class=\"text-lg md:text-xl font-bold text-gray-300 tracking-wide uppercase\">errors</div></div></div></div><!-- Right Column --><div class=\"md:w-4/6 h-full flex\"><div class=\"bg-gray-900 rounded-xl shadow-lg p-8 w-full h-full flex flex-col\"><div class=\"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 grid-rows-4 gap-2 container-group flex-1 items-start justify-items-center\"><!-- Cards will be dynamically inserted here --></div></div></div></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
