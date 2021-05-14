package grob

var TraceTypeStreamtube TraceType = "streamtube"

func (trace *Streamtube) GetType() TraceType {
	return TraceTypeStreamtube
}
// Streamtube Use a streamtube trace to visualize flow in a vector field.  Specify a vector field using 6 1D arrays of equal length, 3 position arrays `x`, `y` and `z` and 3 vector component arrays `u`, `v`, and `w`.  By default, the tubes' starting positions will be cut from the vector field's x-z plane at its minimum y value. To specify your own starting position, use attributes `starts.x`, `starts.y` and `starts.z`. The color is encoded by the norm of (u, v, w), and the local radius by the divergence of (u, v, w).
type Streamtube struct {
    
    // Type 
    // is the type of the plot 
    Type TraceType `json:"type,omitempty"`
    
    // Autocolorscale 
    // boolean 
    // Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed. 
    Autocolorscale Bool `json:"autocolorscale,omitempty"`
    
    // Cauto 
    // boolean 
    // Determines whether or not the color domain is computed with respect to the input data (here u/v/w norm) or the bounds set in `cmin` and `cmax`  Defaults to `false` when `cmin` and `cmax` are set by the user. 
    Cauto Bool `json:"cauto,omitempty"`
    
    // Cmax 
    // number 
    // Sets the upper bound of the color domain. Value should have the same units as u/v/w norm and if set, `cmin` must be set as well. 
    Cmax float64 `json:"cmax,omitempty"`
    
    // Cmid 
    // number 
    // Sets the mid-point of the color domain by scaling `cmin` and/or `cmax` to be equidistant to this point. Value should have the same units as u/v/w norm. Has no effect when `cauto` is `false`. 
    Cmid float64 `json:"cmid,omitempty"`
    
    // Cmin 
    // number 
    // Sets the lower bound of the color domain. Value should have the same units as u/v/w norm and if set, `cmax` must be set as well. 
    Cmin float64 `json:"cmin,omitempty"`
    
    // Coloraxis 
    // subplotid 
    // Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis. 
    Coloraxis String `json:"coloraxis,omitempty"`
    
    // Colorbar 
    //   
    Colorbar *StreamtubeColorbar `json:"colorbar,omitempty"`
    
    // Colorscale 
    // colorscale 
    // Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`cmin` and `cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis. 
    // Default: <nil> 
    Colorscale ColorScale `json:"colorscale,omitempty"`
    
    // Customdata 
    // data_array 
    // Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements 
    Customdata interface{} `json:"customdata,omitempty"`
    
    // Customdatasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  customdata . 
    Customdatasrc String `json:"customdatasrc,omitempty"`
    
    // Hoverinfo 
    // flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. 
    Hoverinfo StreamtubeHoverinfo `json:"hoverinfo,omitempty"`
    
    // Hoverinfosrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hoverinfo . 
    Hoverinfosrc String `json:"hoverinfosrc,omitempty"`
    
    // Hoverlabel 
    //   
    Hoverlabel *StreamtubeHoverlabel `json:"hoverlabel,omitempty"`
    
    // Hovertemplate 
    // string 
    // Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `tubex`, `tubey`, `tubez`, `tubeu`, `tubev`, `tubew`, `norm` and `divergence`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`. 
    Hovertemplate String `json:"hovertemplate,omitempty"`
    
    // Hovertemplatesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hovertemplate . 
    Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`
    
    // Hovertext 
    // string 
    // Same as `text`. 
    Hovertext String `json:"hovertext,omitempty"`
    
    // Ids 
    // data_array 
    // Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type. 
    Ids interface{} `json:"ids,omitempty"`
    
    // Idssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  ids . 
    Idssrc String `json:"idssrc,omitempty"`
    
    // Legendgroup 
    // string 
    // Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items. 
    Legendgroup String `json:"legendgroup,omitempty"`
    
    // Lighting 
    //   
    Lighting *StreamtubeLighting `json:"lighting,omitempty"`
    
    // Lightposition 
    //   
    Lightposition *StreamtubeLightposition `json:"lightposition,omitempty"`
    
    // Maxdisplayed 
    // integer 
    // The maximum number of displayed segments in a streamtube. 
    Maxdisplayed int64 `json:"maxdisplayed,omitempty"`
    
    // Meta 
    // any 
    // Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index. 
    Meta interface{} `json:"meta,omitempty"`
    
    // Metasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  meta . 
    Metasrc String `json:"metasrc,omitempty"`
    
    // Name 
    // string 
    // Sets the trace name. The trace name appear as the legend item and on hover. 
    Name String `json:"name,omitempty"`
    
    // Opacity 
    // number 
    // Sets the opacity of the surface. Please note that in the case of using high `opacity` values for example a value greater than or equal to 0.5 on two surfaces (and 0.25 with four surfaces), an overlay of multiple transparent surfaces may not perfectly be sorted in depth by the webgl API. This behavior may be improved in the near future and is subject to change. 
    Opacity float64 `json:"opacity,omitempty"`
    
    // Reversescale 
    // boolean 
    // Reverses the color mapping if true. If true, `cmin` will correspond to the last color in the array and `cmax` will correspond to the first color. 
    Reversescale Bool `json:"reversescale,omitempty"`
    
    // Scene 
    // subplotid 
    // Sets a reference between this trace's 3D coordinate system and a 3D scene. If *scene* (the default value), the (x,y,z) coordinates refer to `layout.scene`. If *scene2*, the (x,y,z) coordinates refer to `layout.scene2`, and so on. 
    Scene String `json:"scene,omitempty"`
    
    // Showlegend 
    // boolean 
    // Determines whether or not an item corresponding to this trace is shown in the legend. 
    Showlegend Bool `json:"showlegend,omitempty"`
    
    // Showscale 
    // boolean 
    // Determines whether or not a colorbar is displayed for this trace. 
    Showscale Bool `json:"showscale,omitempty"`
    
    // Sizeref 
    // number 
    // The scaling factor for the streamtubes. The default is 1, which avoids two max divergence tubes from touching at adjacent starting positions. 
    Sizeref float64 `json:"sizeref,omitempty"`
    
    // Starts 
    //   
    Starts *StreamtubeStarts `json:"starts,omitempty"`
    
    // Stream 
    //   
    Stream *StreamtubeStream `json:"stream,omitempty"`
    
    // Text 
    // string 
    // Sets a text element associated with this trace. If trace `hoverinfo` contains a *text* flag, this text element will be seen in all hover labels. Note that streamtube traces do not support array `text` values. 
    Text String `json:"text,omitempty"`
    
    // U 
    // data_array 
    // Sets the x components of the vector field. 
    U interface{} `json:"u,omitempty"`
    
    // Uid 
    // string 
    // Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions. 
    Uid String `json:"uid,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Usrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  u . 
    Usrc String `json:"usrc,omitempty"`
    
    // V 
    // data_array 
    // Sets the y components of the vector field. 
    V interface{} `json:"v,omitempty"`
    
    // Visible 
    // enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible). 
    Visible StreamtubeVisible `json:"visible,omitempty"`
    
    // Vsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  v . 
    Vsrc String `json:"vsrc,omitempty"`
    
    // W 
    // data_array 
    // Sets the z components of the vector field. 
    W interface{} `json:"w,omitempty"`
    
    // Wsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  w . 
    Wsrc String `json:"wsrc,omitempty"`
    
    // X 
    // data_array 
    // Sets the x coordinates of the vector field. 
    X interface{} `json:"x,omitempty"`
    
    // Xsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  x . 
    Xsrc String `json:"xsrc,omitempty"`
    
    // Y 
    // data_array 
    // Sets the y coordinates of the vector field. 
    Y interface{} `json:"y,omitempty"`
    
    // Ysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  y . 
    Ysrc String `json:"ysrc,omitempty"`
    
    // Z 
    // data_array 
    // Sets the z coordinates of the vector field. 
    Z interface{} `json:"z,omitempty"`
    
    // Zsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  z . 
    Zsrc String `json:"zsrc,omitempty"`
    
}
// StreamtubeColorbarTickfont Sets the color bar's tick label font
type StreamtubeColorbarTickfont struct {
    
    // Color 
    // color 
    //  
    Color Color `json:"color,omitempty"`
    
    // Family 
    // string 
    // HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*. 
    Family String `json:"family,omitempty"`
    
    // Size 
    // number 
    //  
    Size float64 `json:"size,omitempty"`
    
}
// StreamtubeColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type StreamtubeColorbarTitleFont struct {
    
    // Color 
    // color 
    //  
    Color Color `json:"color,omitempty"`
    
    // Family 
    // string 
    // HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*. 
    Family String `json:"family,omitempty"`
    
    // Size 
    // number 
    //  
    Size float64 `json:"size,omitempty"`
    
}
// StreamtubeColorbarTitle 
type StreamtubeColorbarTitle struct {
    
    // Font 
    //  Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute. 
    Font *StreamtubeColorbarTitleFont `json:"font,omitempty"`
    
    // Side 
    // enumerated Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute. 
    Side StreamtubeColorbarTitleSide `json:"side,omitempty"`
    
    // Text 
    // string 
    // Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// StreamtubeColorbar 
type StreamtubeColorbar struct {
    
    // Bgcolor 
    // color 
    // Sets the color of padded area. 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Bordercolor 
    // color 
    // Sets the axis line color. 
    Bordercolor Color `json:"bordercolor,omitempty"`
    
    // Borderwidth 
    // number 
    // Sets the width (in px) or the border enclosing this color bar. 
    Borderwidth float64 `json:"borderwidth,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat StreamtubeColorbarExponentformat `json:"exponentformat,omitempty"`
    
    // Len 
    // number 
    // Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends. 
    Len float64 `json:"len,omitempty"`
    
    // Lenmode 
    // enumerated Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value. 
    Lenmode StreamtubeColorbarLenmode `json:"lenmode,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Outlinecolor 
    // color 
    // Sets the axis line color. 
    Outlinecolor Color `json:"outlinecolor,omitempty"`
    
    // Outlinewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Outlinewidth float64 `json:"outlinewidth,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent StreamtubeColorbarShowexponent `json:"showexponent,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix StreamtubeColorbarShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix StreamtubeColorbarShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Thickness 
    // number 
    // Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels. 
    Thickness float64 `json:"thickness,omitempty"`
    
    // Thicknessmode 
    // enumerated Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value. 
    Thicknessmode StreamtubeColorbarThicknessmode `json:"thicknessmode,omitempty"`
    
    // Tick0 
    // any 
    // Sets the placement of the first tick on this axis. Use with `dtick`. If the axis `type` is *log*, then you must take the log of your starting tick (e.g. to set the starting tick to 100, set the `tick0` to 2) except when `dtick`=*L<f>* (see `dtick` for more info). If the axis `type` is *date*, it should be a date string, like date data. If the axis `type` is *category*, it should be a number, using the scale where each category is assigned a serial number from zero in the order it appears. 
    Tick0 interface{} `json:"tick0,omitempty"`
    
    // Tickangle 
    // angle 
    // Sets the angle of the tick labels with respect to the horizontal. For example, a `tickangle` of -90 draws the tick labels vertically. 
    Tickangle float64 `json:"tickangle,omitempty"`
    
    // Tickcolor 
    // color 
    // Sets the tick color. 
    Tickcolor Color `json:"tickcolor,omitempty"`
    
    // Tickfont 
    //  Sets the color bar's tick label font 
    Tickfont *StreamtubeColorbarTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklabelposition 
    // enumerated Determines where tick labels are drawn. 
    Ticklabelposition StreamtubeColorbarTicklabelposition `json:"ticklabelposition,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode StreamtubeColorbarTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks StreamtubeColorbarTicks `json:"ticks,omitempty"`
    
    // Ticksuffix 
    // string 
    // Sets a tick label suffix. 
    Ticksuffix String `json:"ticksuffix,omitempty"`
    
    // Ticktext 
    // data_array 
    // Sets the text displayed at the ticks position via `tickvals`. Only has an effect if `tickmode` is set to *array*. Used with `tickvals`. 
    Ticktext interface{} `json:"ticktext,omitempty"`
    
    // Ticktextsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  ticktext . 
    Ticktextsrc String `json:"ticktextsrc,omitempty"`
    
    // Tickvals 
    // data_array 
    // Sets the values at which ticks on this axis appear. Only has an effect if `tickmode` is set to *array*. Used with `ticktext`. 
    Tickvals interface{} `json:"tickvals,omitempty"`
    
    // Tickvalssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  tickvals . 
    Tickvalssrc String `json:"tickvalssrc,omitempty"`
    
    // Tickwidth 
    // number 
    // Sets the tick width (in px). 
    Tickwidth float64 `json:"tickwidth,omitempty"`
    
    // Title 
    //   
    Title *StreamtubeColorbarTitle `json:"title,omitempty"`
    
    // X 
    // number 
    // Sets the x position of the color bar (in plot fraction). 
    X float64 `json:"x,omitempty"`
    
    // Xanchor 
    // enumerated Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar. 
    Xanchor StreamtubeColorbarXanchor `json:"xanchor,omitempty"`
    
    // Xpad 
    // number 
    // Sets the amount of padding (in px) along the x direction. 
    Xpad float64 `json:"xpad,omitempty"`
    
    // Y 
    // number 
    // Sets the y position of the color bar (in plot fraction). 
    Y float64 `json:"y,omitempty"`
    
    // Yanchor 
    // enumerated Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar. 
    Yanchor StreamtubeColorbarYanchor `json:"yanchor,omitempty"`
    
    // Ypad 
    // number 
    // Sets the amount of padding (in px) along the y direction. 
    Ypad float64 `json:"ypad,omitempty"`
    
}
// StreamtubeHoverlabelFont Sets the font used in hover labels.
type StreamtubeHoverlabelFont struct {
    
    // Color 
    // color 
    //  
    Color Color `json:"color,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
    // Family 
    // string 
    // HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*. 
    Family String `json:"family,omitempty"`
    
    // Familysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  family . 
    Familysrc String `json:"familysrc,omitempty"`
    
    // Size 
    // number 
    //  
    Size float64 `json:"size,omitempty"`
    
    // Sizesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  size . 
    Sizesrc String `json:"sizesrc,omitempty"`
    
}
// StreamtubeHoverlabel 
type StreamtubeHoverlabel struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines 
    Align StreamtubeHoverlabelAlign `json:"align,omitempty"`
    
    // Alignsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  align . 
    Alignsrc String `json:"alignsrc,omitempty"`
    
    // Bgcolor 
    // color 
    // Sets the background color of the hover labels for this trace 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Bgcolorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  bgcolor . 
    Bgcolorsrc String `json:"bgcolorsrc,omitempty"`
    
    // Bordercolor 
    // color 
    // Sets the border color of the hover labels for this trace. 
    Bordercolor Color `json:"bordercolor,omitempty"`
    
    // Bordercolorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  bordercolor . 
    Bordercolorsrc String `json:"bordercolorsrc,omitempty"`
    
    // Font 
    //  Sets the font used in hover labels. 
    Font *StreamtubeHoverlabelFont `json:"font,omitempty"`
    
    // Namelength 
    // integer 
    // Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis. 
    Namelength int64 `json:"namelength,omitempty"`
    
    // Namelengthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  namelength . 
    Namelengthsrc String `json:"namelengthsrc,omitempty"`
    
}
// StreamtubeLighting 
type StreamtubeLighting struct {
    
    // Ambient 
    // number 
    // Ambient light increases overall color visibility but can wash out the image. 
    Ambient float64 `json:"ambient,omitempty"`
    
    // Diffuse 
    // number 
    // Represents the extent that incident rays are reflected in a range of angles. 
    Diffuse float64 `json:"diffuse,omitempty"`
    
    // Facenormalsepsilon 
    // number 
    // Epsilon for face normals calculation avoids math issues arising from degenerate geometry. 
    Facenormalsepsilon float64 `json:"facenormalsepsilon,omitempty"`
    
    // Fresnel 
    // number 
    // Represents the reflectance as a dependency of the viewing angle; e.g. paper is reflective when viewing it from the edge of the paper (almost 90 degrees), causing shine. 
    Fresnel float64 `json:"fresnel,omitempty"`
    
    // Roughness 
    // number 
    // Alters specular reflection; the rougher the surface, the wider and less contrasty the shine. 
    Roughness float64 `json:"roughness,omitempty"`
    
    // Specular 
    // number 
    // Represents the level that incident rays are reflected in a single direction, causing shine. 
    Specular float64 `json:"specular,omitempty"`
    
    // Vertexnormalsepsilon 
    // number 
    // Epsilon for vertex normals calculation avoids math issues arising from degenerate geometry. 
    Vertexnormalsepsilon float64 `json:"vertexnormalsepsilon,omitempty"`
    
}
// StreamtubeLightposition 
type StreamtubeLightposition struct {
    
    // X 
    // number 
    // Numeric vector, representing the X coordinate for each vertex. 
    X float64 `json:"x,omitempty"`
    
    // Y 
    // number 
    // Numeric vector, representing the Y coordinate for each vertex. 
    Y float64 `json:"y,omitempty"`
    
    // Z 
    // number 
    // Numeric vector, representing the Z coordinate for each vertex. 
    Z float64 `json:"z,omitempty"`
    
}
// StreamtubeStarts 
type StreamtubeStarts struct {
    
    // X 
    // data_array 
    // Sets the x components of the starting position of the streamtubes 
    X interface{} `json:"x,omitempty"`
    
    // Xsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  x . 
    Xsrc String `json:"xsrc,omitempty"`
    
    // Y 
    // data_array 
    // Sets the y components of the starting position of the streamtubes 
    Y interface{} `json:"y,omitempty"`
    
    // Ysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  y . 
    Ysrc String `json:"ysrc,omitempty"`
    
    // Z 
    // data_array 
    // Sets the z components of the starting position of the streamtubes 
    Z interface{} `json:"z,omitempty"`
    
    // Zsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  z . 
    Zsrc String `json:"zsrc,omitempty"`
    
}
// StreamtubeStream 
type StreamtubeStream struct {
    
    // Maxpoints 
    // number 
    // Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot. 
    Maxpoints float64 `json:"maxpoints,omitempty"`
    
    // Token 
    // string 
    // The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details. 
    Token String `json:"token,omitempty"`
    
}
// StreamtubeColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type StreamtubeColorbarExponentformat string 

const (
    StreamtubeColorbarExponentformatNone StreamtubeColorbarExponentformat = "none"
    StreamtubeColorbarExponentformatE1 StreamtubeColorbarExponentformat = "e"
    StreamtubeColorbarExponentformatE2 StreamtubeColorbarExponentformat = "E"
    StreamtubeColorbarExponentformatPower StreamtubeColorbarExponentformat = "power"
    StreamtubeColorbarExponentformatSi StreamtubeColorbarExponentformat = "SI"
    StreamtubeColorbarExponentformatB StreamtubeColorbarExponentformat = "B"
    
)
// StreamtubeColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type StreamtubeColorbarLenmode string 

const (
    StreamtubeColorbarLenmodeFraction StreamtubeColorbarLenmode = "fraction"
    StreamtubeColorbarLenmodePixels StreamtubeColorbarLenmode = "pixels"
    
)
// StreamtubeColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type StreamtubeColorbarShowexponent string 

const (
    StreamtubeColorbarShowexponentAll StreamtubeColorbarShowexponent = "all"
    StreamtubeColorbarShowexponentFirst StreamtubeColorbarShowexponent = "first"
    StreamtubeColorbarShowexponentLast StreamtubeColorbarShowexponent = "last"
    StreamtubeColorbarShowexponentNone StreamtubeColorbarShowexponent = "none"
    
)
// StreamtubeColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type StreamtubeColorbarShowtickprefix string 

const (
    StreamtubeColorbarShowtickprefixAll StreamtubeColorbarShowtickprefix = "all"
    StreamtubeColorbarShowtickprefixFirst StreamtubeColorbarShowtickprefix = "first"
    StreamtubeColorbarShowtickprefixLast StreamtubeColorbarShowtickprefix = "last"
    StreamtubeColorbarShowtickprefixNone StreamtubeColorbarShowtickprefix = "none"
    
)
// StreamtubeColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type StreamtubeColorbarShowticksuffix string 

const (
    StreamtubeColorbarShowticksuffixAll StreamtubeColorbarShowticksuffix = "all"
    StreamtubeColorbarShowticksuffixFirst StreamtubeColorbarShowticksuffix = "first"
    StreamtubeColorbarShowticksuffixLast StreamtubeColorbarShowticksuffix = "last"
    StreamtubeColorbarShowticksuffixNone StreamtubeColorbarShowticksuffix = "none"
    
)
// StreamtubeColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type StreamtubeColorbarThicknessmode string 

const (
    StreamtubeColorbarThicknessmodeFraction StreamtubeColorbarThicknessmode = "fraction"
    StreamtubeColorbarThicknessmodePixels StreamtubeColorbarThicknessmode = "pixels"
    
)
// StreamtubeColorbarTicklabelposition Determines where tick labels are drawn.
type StreamtubeColorbarTicklabelposition string 

const (
    StreamtubeColorbarTicklabelpositionOutside StreamtubeColorbarTicklabelposition = "outside"
    StreamtubeColorbarTicklabelpositionInside StreamtubeColorbarTicklabelposition = "inside"
    StreamtubeColorbarTicklabelpositionOutsideTop StreamtubeColorbarTicklabelposition = "outside top"
    StreamtubeColorbarTicklabelpositionInsideTop StreamtubeColorbarTicklabelposition = "inside top"
    StreamtubeColorbarTicklabelpositionOutsideBottom StreamtubeColorbarTicklabelposition = "outside bottom"
    StreamtubeColorbarTicklabelpositionInsideBottom StreamtubeColorbarTicklabelposition = "inside bottom"
    
)
// StreamtubeColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type StreamtubeColorbarTickmode string 

const (
    StreamtubeColorbarTickmodeAuto StreamtubeColorbarTickmode = "auto"
    StreamtubeColorbarTickmodeLinear StreamtubeColorbarTickmode = "linear"
    StreamtubeColorbarTickmodeArray StreamtubeColorbarTickmode = "array"
    
)
// StreamtubeColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type StreamtubeColorbarTicks string 

const (
    StreamtubeColorbarTicksOutside StreamtubeColorbarTicks = "outside"
    StreamtubeColorbarTicksInside StreamtubeColorbarTicks = "inside"
    StreamtubeColorbarTicksEmpty StreamtubeColorbarTicks = ""
    
)
// StreamtubeColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type StreamtubeColorbarTitleSide string 

const (
    StreamtubeColorbarTitleSideRight StreamtubeColorbarTitleSide = "right"
    StreamtubeColorbarTitleSideTop StreamtubeColorbarTitleSide = "top"
    StreamtubeColorbarTitleSideBottom StreamtubeColorbarTitleSide = "bottom"
    
)
// StreamtubeColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type StreamtubeColorbarXanchor string 

const (
    StreamtubeColorbarXanchorLeft StreamtubeColorbarXanchor = "left"
    StreamtubeColorbarXanchorCenter StreamtubeColorbarXanchor = "center"
    StreamtubeColorbarXanchorRight StreamtubeColorbarXanchor = "right"
    
)
// StreamtubeColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type StreamtubeColorbarYanchor string 

const (
    StreamtubeColorbarYanchorTop StreamtubeColorbarYanchor = "top"
    StreamtubeColorbarYanchorMiddle StreamtubeColorbarYanchor = "middle"
    StreamtubeColorbarYanchorBottom StreamtubeColorbarYanchor = "bottom"
    
)
// StreamtubeHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type StreamtubeHoverlabelAlign string 

const (
    StreamtubeHoverlabelAlignLeft StreamtubeHoverlabelAlign = "left"
    StreamtubeHoverlabelAlignRight StreamtubeHoverlabelAlign = "right"
    StreamtubeHoverlabelAlignAuto StreamtubeHoverlabelAlign = "auto"
    
)
// StreamtubeVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type StreamtubeVisible interface{} 

var (
    StreamtubeVisibleTrue StreamtubeVisible = true
    StreamtubeVisibleFalse StreamtubeVisible = false
    StreamtubeVisibleLegendonly StreamtubeVisible = "legendonly"
    
)
// StreamtubeHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type StreamtubeHoverinfo string 

const (
    // Flags
    StreamtubeHoverinfoX StreamtubeHoverinfo = "x"
    StreamtubeHoverinfoY StreamtubeHoverinfo = "y"
    StreamtubeHoverinfoZ StreamtubeHoverinfo = "z"
    StreamtubeHoverinfoU StreamtubeHoverinfo = "u"
    StreamtubeHoverinfoV StreamtubeHoverinfo = "v"
    StreamtubeHoverinfoW StreamtubeHoverinfo = "w"
    StreamtubeHoverinfoNorm StreamtubeHoverinfo = "norm"
    StreamtubeHoverinfoDivergence StreamtubeHoverinfo = "divergence"
    StreamtubeHoverinfoText StreamtubeHoverinfo = "text"
    StreamtubeHoverinfoName StreamtubeHoverinfo = "name"
    
    // Extra
    StreamtubeHoverinfoAll StreamtubeHoverinfo = "all"
    StreamtubeHoverinfoNone StreamtubeHoverinfo = "none"
    StreamtubeHoverinfoSkip StreamtubeHoverinfo = "skip"
    
)