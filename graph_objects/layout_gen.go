package grob

// Layout Plot layout options
type Layout struct {
    
    // Activeshape 
    //   
    Activeshape *LayoutActiveshape `json:"activeshape,omitempty"`
    
    // Angularaxis 
    //   
    Angularaxis *LayoutAngularaxis `json:"angularaxis,omitempty"`
    
    // Annotations 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Annotations interface{} `json:"annotations,omitempty"`
    
    // Autosize 
    // boolean 
    // Determines whether or not a layout width or height that has been left undefined by the user is initialized on each relayout. Note that, regardless of this attribute, an undefined layout width or height is always initialized on the first call to plot. 
    Autosize Bool `json:"autosize,omitempty"`
    
    // Autotypenumbers 
    // enumerated Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. This is the default value; however it could be overridden for individual axes. 
    Autotypenumbers LayoutAutotypenumbers `json:"autotypenumbers,omitempty"`
    
    // Calendar 
    // enumerated Sets the default calendar system to use for interpreting and displaying dates throughout the plot. 
    Calendar LayoutCalendar `json:"calendar,omitempty"`
    
    // Clickmode 
    // flaglist Determines the mode of single click interactions. *event* is the default value and emits the `plotly_click` event. In addition this mode emits the `plotly_selected` event in drag modes *lasso* and *select*, but with no event data attached (kept for compatibility reasons). The *select* flag enables selecting single data points via click. This mode also supports persistent selections, meaning that pressing Shift while clicking, adds to / subtracts from an existing selection. *select* with `hovermode`: *x* can be confusing, consider explicitly setting `hovermode`: *closest* when using this feature. Selection events are sent accordingly as long as *event* flag is set as well. When the *event* flag is missing, `plotly_click` and `plotly_selected` events are not fired. 
    Clickmode LayoutClickmode `json:"clickmode,omitempty"`
    
    // Coloraxis 
    //   
    Coloraxis *LayoutColoraxis `json:"coloraxis,omitempty"`
    
    // Colorscale 
    //   
    Colorscale *LayoutColorscale `json:"colorscale,omitempty"`
    
    // Colorway 
    // colorlist 
    // Sets the default trace colors. 
    Colorway ColorList `json:"colorway,omitempty"`
    
    // Computed 
    // any 
    // Placeholder for exporting automargin-impacting values namely `margin.t`, `margin.b`, `margin.l` and `margin.r` in *full-json* mode. 
    Computed interface{} `json:"computed,omitempty"`
    
    // Datarevision 
    // any 
    // If provided, a changed value tells `Plotly.react` that one or more data arrays has changed. This way you can modify arrays in-place rather than making a complete new copy for an incremental change. If NOT provided, `Plotly.react` assumes that data arrays are being treated as immutable, thus any data array with a different identity from its predecessor contains new data. 
    Datarevision interface{} `json:"datarevision,omitempty"`
    
    // Direction 
    // enumerated Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the direction corresponding to positive angles in legacy polar charts. 
    Direction LayoutDirection `json:"direction,omitempty"`
    
    // Dragmode 
    // enumerated Determines the mode of drag interactions. *select* and *lasso* apply only to scatter traces with markers or text. *orbit* and *turntable* apply only to 3D scenes. 
    Dragmode LayoutDragmode `json:"dragmode,omitempty"`
    
    // Editrevision 
    // any 
    // Controls persistence of user-driven changes in `editable: true` configuration, other than trace names and axis titles. Defaults to `layout.uirevision`. 
    Editrevision interface{} `json:"editrevision,omitempty"`
    
    // Font 
    //  Sets the global font. Note that fonts used in traces and other layout components inherit from the global font. 
    Font *LayoutFont `json:"font,omitempty"`
    
    // Geo 
    //   
    Geo *LayoutGeo `json:"geo,omitempty"`
    
    // Grid 
    //   
    Grid *LayoutGrid `json:"grid,omitempty"`
    
    // Height 
    // number 
    // Sets the plot's height (in px). 
    Height float64 `json:"height,omitempty"`
    
    // Hidesources 
    // boolean 
    // Determines whether or not a text link citing the data source is placed at the bottom-right cored of the figure. Has only an effect only on graphs that have been generated via forked graphs from the Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise). 
    Hidesources Bool `json:"hidesources,omitempty"`
    
    // Hoverdistance 
    // integer 
    // Sets the default distance (in pixels) to look for data to add hover labels (-1 means no cutoff, 0 means no looking for data). This is only a real distance for hovering on point-like objects, like scatter points. For area-like objects (bars, scatter fills, etc) hovering is on inside the area and off outside, but these objects will not supersede hover on point-like objects in case of conflict. 
    Hoverdistance int64 `json:"hoverdistance,omitempty"`
    
    // Hoverlabel 
    //   
    Hoverlabel *LayoutHoverlabel `json:"hoverlabel,omitempty"`
    
    // Hovermode 
    // enumerated Determines the mode of hover interactions. If *closest*, a single hoverlabel will appear for the *closest* point within the `hoverdistance`. If *x* (or *y*), multiple hoverlabels will appear for multiple points at the *closest* x- (or y-) coordinate within the `hoverdistance`, with the caveat that no more than one hoverlabel will appear per trace. If *x unified* (or *y unified*), a single hoverlabel will appear multiple points at the closest x- (or y-) coordinate within the `hoverdistance` with the caveat that no more than one hoverlabel will appear per trace. In this mode, spikelines are enabled by default perpendicular to the specified axis. If false, hover interactions are disabled. If `clickmode` includes the *select* flag, `hovermode` defaults to *closest*. If `clickmode` lacks the *select* flag, it defaults to *x* or *y* (depending on the trace's `orientation` value) for plots based on cartesian coordinates. For anything else the default value is *closest*. 
    Hovermode LayoutHovermode `json:"hovermode,omitempty"`
    
    // Images 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Images interface{} `json:"images,omitempty"`
    
    // Legend 
    //   
    Legend *LayoutLegend `json:"legend,omitempty"`
    
    // Mapbox 
    //   
    Mapbox *LayoutMapbox `json:"mapbox,omitempty"`
    
    // Margin 
    //   
    Margin *LayoutMargin `json:"margin,omitempty"`
    
    // Meta 
    // any 
    // Assigns extra meta information that can be used in various `text` attributes. Attributes such as the graph, axis and colorbar `title.text`, annotation `text` `trace.name` in legend items, `rangeselector`, `updatemenus` and `sliders` `label` text all support `meta`. One can access `meta` fields using template strings: `%{meta[i]}` where `i` is the index of the `meta` item in question. `meta` can also be an object for example `{key: value}` which can be accessed %{meta[key]}. 
    Meta interface{} `json:"meta,omitempty"`
    
    // Metasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  meta . 
    Metasrc String `json:"metasrc,omitempty"`
    
    // Modebar 
    //   
    Modebar *LayoutModebar `json:"modebar,omitempty"`
    
    // Newshape 
    //   
    Newshape *LayoutNewshape `json:"newshape,omitempty"`
    
    // Orientation 
    // angle 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Rotates the entire polar by the given angle in legacy polar charts. 
    Orientation float64 `json:"orientation,omitempty"`
    
    // PaperBgcolor 
    // color 
    // Sets the background color of the paper where the graph is drawn. 
    PaperBgcolor Color `json:"paper_bgcolor,omitempty"`
    
    // PlotBgcolor 
    // color 
    // Sets the background color of the plotting area in-between x and y axes. 
    PlotBgcolor Color `json:"plot_bgcolor,omitempty"`
    
    // Polar 
    //   
    Polar *LayoutPolar `json:"polar,omitempty"`
    
    // Radialaxis 
    //   
    Radialaxis *LayoutRadialaxis `json:"radialaxis,omitempty"`
    
    // Scene 
    //   
    Scene *LayoutScene `json:"scene,omitempty"`
    
    // Selectdirection 
    // enumerated When `dragmode` is set to *select*, this limits the selection of the drag to horizontal, vertical or diagonal. *h* only allows horizontal selection, *v* only vertical, *d* only diagonal and *any* sets no limit. 
    Selectdirection LayoutSelectdirection `json:"selectdirection,omitempty"`
    
    // Selectionrevision 
    // any 
    // Controls persistence of user-driven changes in selected points from all traces. 
    Selectionrevision interface{} `json:"selectionrevision,omitempty"`
    
    // Separators 
    // string 
    // Sets the decimal and thousand separators. For example, *. * puts a '.' before decimals and a space between thousands. In English locales, dflt is *.,* but other locales may alter this default. 
    Separators String `json:"separators,omitempty"`
    
    // Shapes 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Shapes interface{} `json:"shapes,omitempty"`
    
    // Showlegend 
    // boolean 
    // Determines whether or not a legend is drawn. Default is `true` if there is a trace to show and any of these: a) Two or more traces would by default be shown in the legend. b) One pie trace is shown in the legend. c) One trace is explicitly given with `showlegend: true`. 
    Showlegend Bool `json:"showlegend,omitempty"`
    
    // Sliders 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Sliders interface{} `json:"sliders,omitempty"`
    
    // Spikedistance 
    // integer 
    // Sets the default distance (in pixels) to look for data to draw spikelines to (-1 means no cutoff, 0 means no looking for data). As with hoverdistance, distance does not apply to area-like objects. In addition, some objects can be hovered on but will not generate spikelines, such as scatter fills. 
    Spikedistance int64 `json:"spikedistance,omitempty"`
    
    // Template 
    // any 
    // Default attributes to be applied to the plot. Templates can be created from existing plots using `Plotly.makeTemplate`, or created manually. They should be objects with format: `{layout: layoutTemplate, data: {[type]: [traceTemplate, ...]}, ...}` `layoutTemplate` and `traceTemplate` are objects matching the attribute structure of `layout` and a data trace.  Trace templates are applied cyclically to traces of each type. Container arrays (eg `annotations`) have special handling: An object ending in `defaults` (eg `annotationdefaults`) is applied to each array item. But if an item has a `templateitemname` key we look in the template array for an item with matching `name` and apply that instead. If no matching `name` is found we mark the item invisible. Any named template item not referenced is appended to the end of the array, so you can use this for a watermark annotation or a logo image, for example. To omit one of these items on the plot, make an item with matching `templateitemname` and `visible: false`. 
    Template interface{} `json:"template,omitempty"`
    
    // Ternary 
    //   
    Ternary *LayoutTernary `json:"ternary,omitempty"`
    
    // Title 
    //   
    Title *LayoutTitle `json:"title,omitempty"`
    
    // Transition 
    //  Sets transition options used during Plotly.react updates. 
    Transition *LayoutTransition `json:"transition,omitempty"`
    
    // Uirevision 
    // any 
    // Used to allow user interactions with the plot to persist after `Plotly.react` calls that are unaware of these interactions. If `uirevision` is omitted, or if it is given and it changed from the previous `Plotly.react` call, the exact new figure is used. If `uirevision` is truthy and did NOT change, any attribute that has been affected by user interactions and did not receive a different value in the new figure will keep the interaction value. `layout.uirevision` attribute serves as the default for `uirevision` attributes in various sub-containers. For finer control you can set these sub-attributes directly. For example, if your app separately controls the data on the x and y axes you might set `xaxis.uirevision=*time*` and `yaxis.uirevision=*cost*`. Then if only the y data is changed, you can update `yaxis.uirevision=*quantity*` and the y axis range will reset but the x axis range will retain any user-driven zoom. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Uniformtext 
    //   
    Uniformtext *LayoutUniformtext `json:"uniformtext,omitempty"`
    
    // Updatemenus 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Updatemenus interface{} `json:"updatemenus,omitempty"`
    
    // Width 
    // number 
    // Sets the plot's width (in px). 
    Width float64 `json:"width,omitempty"`
    
    // Xaxis 
    //   
    Xaxis *LayoutXaxis `json:"xaxis,omitempty"`
    
    // Yaxis 
    //   
    Yaxis *LayoutYaxis `json:"yaxis,omitempty"`
    
    // Bargap 
    // number 
    // Sets the gap (in plot fraction) between bars of adjacent location coordinates. 
    Bargap float64 `json:"bargap,omitempty"`
    
    // Bargroupgap 
    // number 
    // Sets the gap (in plot fraction) between bars of the same location coordinate. 
    Bargroupgap float64 `json:"bargroupgap,omitempty"`
    
    // Barmode 
    // enumerated Determines how bars at the same location coordinate are displayed on the graph. With *stack*, the bars are stacked on top of one another With *relative*, the bars are stacked on top of one another, with negative values below the axis, positive values above With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars. 
    Barmode LayoutBarmode `json:"barmode,omitempty"`
    
    // Barnorm 
    // enumerated Sets the normalization for bar traces on the graph. With *fraction*, the value of each bar is divided by the sum of all values at that location coordinate. *percent* is the same but multiplied by 100 to show percentages. 
    Barnorm LayoutBarnorm `json:"barnorm,omitempty"`
    
    // Boxgap 
    // number 
    // Sets the gap (in plot fraction) between boxes of adjacent location coordinates. Has no effect on traces that have *width* set. 
    Boxgap float64 `json:"boxgap,omitempty"`
    
    // Boxgroupgap 
    // number 
    // Sets the gap (in plot fraction) between boxes of the same location coordinate. Has no effect on traces that have *width* set. 
    Boxgroupgap float64 `json:"boxgroupgap,omitempty"`
    
    // Boxmode 
    // enumerated Determines how boxes at the same location coordinate are displayed on the graph. If *group*, the boxes are plotted next to one another centered around the shared location. If *overlay*, the boxes are plotted over one another, you might need to set *opacity* to see them multiple boxes. Has no effect on traces that have *width* set. 
    Boxmode LayoutBoxmode `json:"boxmode,omitempty"`
    
    // Waterfallgap 
    // number 
    // Sets the gap (in plot fraction) between bars of adjacent location coordinates. 
    Waterfallgap float64 `json:"waterfallgap,omitempty"`
    
    // Waterfallgroupgap 
    // number 
    // Sets the gap (in plot fraction) between bars of the same location coordinate. 
    Waterfallgroupgap float64 `json:"waterfallgroupgap,omitempty"`
    
    // Waterfallmode 
    // enumerated Determines how bars at the same location coordinate are displayed on the graph. With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars. 
    Waterfallmode LayoutWaterfallmode `json:"waterfallmode,omitempty"`
    
    // Extendsunburstcolors 
    // boolean 
    // If `true`, the sunburst slice colors (whether given by `sunburstcolorway` or inherited from `colorway`) will be extended to three times its original length by first repeating every color 20% lighter then each color 20% darker. This is intended to reduce the likelihood of reusing the same color when you have many slices, but you can set `false` to disable. Colors provided in the trace, using `marker.colors`, are never extended. 
    Extendsunburstcolors Bool `json:"extendsunburstcolors,omitempty"`
    
    // Sunburstcolorway 
    // colorlist 
    // Sets the default sunburst slice colors. Defaults to the main `colorway` used for trace colors. If you specify a new list here it can still be extended with lighter and darker colors, see `extendsunburstcolors`. 
    Sunburstcolorway ColorList `json:"sunburstcolorway,omitempty"`
    
    // Extendfunnelareacolors 
    // boolean 
    // If `true`, the funnelarea slice colors (whether given by `funnelareacolorway` or inherited from `colorway`) will be extended to three times its original length by first repeating every color 20% lighter then each color 20% darker. This is intended to reduce the likelihood of reusing the same color when you have many slices, but you can set `false` to disable. Colors provided in the trace, using `marker.colors`, are never extended. 
    Extendfunnelareacolors Bool `json:"extendfunnelareacolors,omitempty"`
    
    // Funnelareacolorway 
    // colorlist 
    // Sets the default funnelarea slice colors. Defaults to the main `colorway` used for trace colors. If you specify a new list here it can still be extended with lighter and darker colors, see `extendfunnelareacolors`. 
    Funnelareacolorway ColorList `json:"funnelareacolorway,omitempty"`
    
    // Hiddenlabels 
    // data_array 
    // hiddenlabels is the funnelarea & pie chart analog of visible:'legendonly' but it can contain many labels, and can simultaneously hide slices from several pies/funnelarea charts 
    Hiddenlabels interface{} `json:"hiddenlabels,omitempty"`
    
    // Hiddenlabelssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hiddenlabels . 
    Hiddenlabelssrc String `json:"hiddenlabelssrc,omitempty"`
    
    // Violingap 
    // number 
    // Sets the gap (in plot fraction) between violins of adjacent location coordinates. Has no effect on traces that have *width* set. 
    Violingap float64 `json:"violingap,omitempty"`
    
    // Violingroupgap 
    // number 
    // Sets the gap (in plot fraction) between violins of the same location coordinate. Has no effect on traces that have *width* set. 
    Violingroupgap float64 `json:"violingroupgap,omitempty"`
    
    // Violinmode 
    // enumerated Determines how violins at the same location coordinate are displayed on the graph. If *group*, the violins are plotted next to one another centered around the shared location. If *overlay*, the violins are plotted over one another, you might need to set *opacity* to see them multiple violins. Has no effect on traces that have *width* set. 
    Violinmode LayoutViolinmode `json:"violinmode,omitempty"`
    
    // Funnelgap 
    // number 
    // Sets the gap (in plot fraction) between bars of adjacent location coordinates. 
    Funnelgap float64 `json:"funnelgap,omitempty"`
    
    // Funnelgroupgap 
    // number 
    // Sets the gap (in plot fraction) between bars of the same location coordinate. 
    Funnelgroupgap float64 `json:"funnelgroupgap,omitempty"`
    
    // Funnelmode 
    // enumerated Determines how bars at the same location coordinate are displayed on the graph. With *stack*, the bars are stacked on top of one another With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars. 
    Funnelmode LayoutFunnelmode `json:"funnelmode,omitempty"`
    
    // Extendpiecolors 
    // boolean 
    // If `true`, the pie slice colors (whether given by `piecolorway` or inherited from `colorway`) will be extended to three times its original length by first repeating every color 20% lighter then each color 20% darker. This is intended to reduce the likelihood of reusing the same color when you have many slices, but you can set `false` to disable. Colors provided in the trace, using `marker.colors`, are never extended. 
    Extendpiecolors Bool `json:"extendpiecolors,omitempty"`
    
    // Piecolorway 
    // colorlist 
    // Sets the default pie slice colors. Defaults to the main `colorway` used for trace colors. If you specify a new list here it can still be extended with lighter and darker colors, see `extendpiecolors`. 
    Piecolorway ColorList `json:"piecolorway,omitempty"`
    
    // Extendtreemapcolors 
    // boolean 
    // If `true`, the treemap slice colors (whether given by `treemapcolorway` or inherited from `colorway`) will be extended to three times its original length by first repeating every color 20% lighter then each color 20% darker. This is intended to reduce the likelihood of reusing the same color when you have many slices, but you can set `false` to disable. Colors provided in the trace, using `marker.colors`, are never extended. 
    Extendtreemapcolors Bool `json:"extendtreemapcolors,omitempty"`
    
    // Treemapcolorway 
    // colorlist 
    // Sets the default treemap slice colors. Defaults to the main `colorway` used for trace colors. If you specify a new list here it can still be extended with lighter and darker colors, see `extendtreemapcolors`. 
    Treemapcolorway ColorList `json:"treemapcolorway,omitempty"`
    
}
// LayoutActiveshape 
type LayoutActiveshape struct {
    
    // Fillcolor 
    // color 
    // Sets the color filling the active shape' interior. 
    Fillcolor Color `json:"fillcolor,omitempty"`
    
    // Opacity 
    // number 
    // Sets the opacity of the active shape. 
    Opacity float64 `json:"opacity,omitempty"`
    
}
// LayoutAngularaxis 
type LayoutAngularaxis struct {
    
    // Domain 
    // info_array 
    // Polar chart subplots are not supported yet. This key has currently no effect. 
    Domain interface{} `json:"domain,omitempty"`
    
    // Endpadding 
    // number 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. 
    Endpadding float64 `json:"endpadding,omitempty"`
    
    // Range 
    // info_array 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Defines the start and end point of this angular axis. 
    Range interface{} `json:"range,omitempty"`
    
    // Showline 
    // boolean 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Determines whether or not the line bounding this angular axis will be shown on the figure. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Determines whether or not the angular axis ticks will feature tick labels. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Tickcolor 
    // color 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the color of the tick lines on this angular axis. 
    Tickcolor Color `json:"tickcolor,omitempty"`
    
    // Ticklen 
    // number 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the length of the tick lines on this angular axis. 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickorientation 
    // enumerated Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the orientation (from the paper perspective) of the angular axis tick labels. 
    Tickorientation LayoutAngularaxisTickorientation `json:"tickorientation,omitempty"`
    
    // Ticksuffix 
    // string 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the length of the tick lines on this angular axis. 
    Ticksuffix String `json:"ticksuffix,omitempty"`
    
    // Visible 
    // boolean 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Determines whether or not this axis will be visible. 
    Visible Bool `json:"visible,omitempty"`
    
}
// LayoutColoraxisColorbarTickfont Sets the color bar's tick label font
type LayoutColoraxisColorbarTickfont struct {
    
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
// LayoutColoraxisColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type LayoutColoraxisColorbarTitleFont struct {
    
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
// LayoutColoraxisColorbarTitle 
type LayoutColoraxisColorbarTitle struct {
    
    // Font 
    //  Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute. 
    Font *LayoutColoraxisColorbarTitleFont `json:"font,omitempty"`
    
    // Side 
    // enumerated Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute. 
    Side LayoutColoraxisColorbarTitleSide `json:"side,omitempty"`
    
    // Text 
    // string 
    // Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// LayoutColoraxisColorbar 
type LayoutColoraxisColorbar struct {
    
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
    Exponentformat LayoutColoraxisColorbarExponentformat `json:"exponentformat,omitempty"`
    
    // Len 
    // number 
    // Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends. 
    Len float64 `json:"len,omitempty"`
    
    // Lenmode 
    // enumerated Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value. 
    Lenmode LayoutColoraxisColorbarLenmode `json:"lenmode,omitempty"`
    
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
    Showexponent LayoutColoraxisColorbarShowexponent `json:"showexponent,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutColoraxisColorbarShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutColoraxisColorbarShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Thickness 
    // number 
    // Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels. 
    Thickness float64 `json:"thickness,omitempty"`
    
    // Thicknessmode 
    // enumerated Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value. 
    Thicknessmode LayoutColoraxisColorbarThicknessmode `json:"thicknessmode,omitempty"`
    
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
    Tickfont *LayoutColoraxisColorbarTickfont `json:"tickfont,omitempty"`
    
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
    Ticklabelposition LayoutColoraxisColorbarTicklabelposition `json:"ticklabelposition,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutColoraxisColorbarTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutColoraxisColorbarTicks `json:"ticks,omitempty"`
    
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
    Title *LayoutColoraxisColorbarTitle `json:"title,omitempty"`
    
    // X 
    // number 
    // Sets the x position of the color bar (in plot fraction). 
    X float64 `json:"x,omitempty"`
    
    // Xanchor 
    // enumerated Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar. 
    Xanchor LayoutColoraxisColorbarXanchor `json:"xanchor,omitempty"`
    
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
    Yanchor LayoutColoraxisColorbarYanchor `json:"yanchor,omitempty"`
    
    // Ypad 
    // number 
    // Sets the amount of padding (in px) along the y direction. 
    Ypad float64 `json:"ypad,omitempty"`
    
}
// LayoutColoraxis 
type LayoutColoraxis struct {
    
    // Autocolorscale 
    // boolean 
    // Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed. 
    Autocolorscale Bool `json:"autocolorscale,omitempty"`
    
    // Cauto 
    // boolean 
    // Determines whether or not the color domain is computed with respect to the input data (here corresponding trace color array(s)) or the bounds set in `cmin` and `cmax`  Defaults to `false` when `cmin` and `cmax` are set by the user. 
    Cauto Bool `json:"cauto,omitempty"`
    
    // Cmax 
    // number 
    // Sets the upper bound of the color domain. Value should have the same units as corresponding trace color array(s) and if set, `cmin` must be set as well. 
    Cmax float64 `json:"cmax,omitempty"`
    
    // Cmid 
    // number 
    // Sets the mid-point of the color domain by scaling `cmin` and/or `cmax` to be equidistant to this point. Value should have the same units as corresponding trace color array(s). Has no effect when `cauto` is `false`. 
    Cmid float64 `json:"cmid,omitempty"`
    
    // Cmin 
    // number 
    // Sets the lower bound of the color domain. Value should have the same units as corresponding trace color array(s) and if set, `cmax` must be set as well. 
    Cmin float64 `json:"cmin,omitempty"`
    
    // Colorbar 
    //   
    Colorbar *LayoutColoraxisColorbar `json:"colorbar,omitempty"`
    
    // Colorscale 
    // colorscale 
    // Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`cmin` and `cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis. 
    // Default: <nil> 
    Colorscale ColorScale `json:"colorscale,omitempty"`
    
    // Reversescale 
    // boolean 
    // Reverses the color mapping if true. If true, `cmin` will correspond to the last color in the array and `cmax` will correspond to the first color. 
    Reversescale Bool `json:"reversescale,omitempty"`
    
    // Showscale 
    // boolean 
    // Determines whether or not a colorbar is displayed for this trace. 
    Showscale Bool `json:"showscale,omitempty"`
    
}
// LayoutColorscale 
type LayoutColorscale struct {
    
    // Diverging 
    // colorscale 
    // Sets the default diverging colorscale. Note that `autocolorscale` must be true for this attribute to work. 
    // Default: [[0 rgb(5,10,172)] [0.35 rgb(106,137,247)] [0.5 rgb(190,190,190)] [0.6 rgb(220,170,132)] [0.7 rgb(230,145,90)] [1 rgb(178,10,28)]] 
    Diverging ColorScale `json:"diverging,omitempty"`
    
    // Sequential 
    // colorscale 
    // Sets the default sequential colorscale for positive values. Note that `autocolorscale` must be true for this attribute to work. 
    // Default: [[0 rgb(220,220,220)] [0.2 rgb(245,195,157)] [0.4 rgb(245,160,105)] [1 rgb(178,10,28)]] 
    Sequential ColorScale `json:"sequential,omitempty"`
    
    // Sequentialminus 
    // colorscale 
    // Sets the default sequential colorscale for negative values. Note that `autocolorscale` must be true for this attribute to work. 
    // Default: [[0 rgb(5,10,172)] [0.35 rgb(40,60,190)] [0.5 rgb(70,100,245)] [0.6 rgb(90,120,245)] [0.7 rgb(106,137,247)] [1 rgb(220,220,220)]] 
    Sequentialminus ColorScale `json:"sequentialminus,omitempty"`
    
}
// LayoutFont Sets the global font. Note that fonts used in traces and other layout components inherit from the global font.
type LayoutFont struct {
    
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
// LayoutGeoCenter 
type LayoutGeoCenter struct {
    
    // Lat 
    // number 
    // Sets the latitude of the map's center. For all projection types, the map's latitude center lies at the middle of the latitude range by default. 
    Lat float64 `json:"lat,omitempty"`
    
    // Lon 
    // number 
    // Sets the longitude of the map's center. By default, the map's longitude center lies at the middle of the longitude range for scoped projection and above `projection.rotation.lon` otherwise. 
    Lon float64 `json:"lon,omitempty"`
    
}
// LayoutGeoDomain 
type LayoutGeoDomain struct {
    
    // Column 
    // integer 
    // If there is a layout grid, use the domain for this column in the grid for this geo subplot . Note that geo subplots are constrained by domain. In general, when `projection.scale` is set to 1. a map will fit either its x or y domain, but not both. 
    Column int64 `json:"column,omitempty"`
    
    // Row 
    // integer 
    // If there is a layout grid, use the domain for this row in the grid for this geo subplot . Note that geo subplots are constrained by domain. In general, when `projection.scale` is set to 1. a map will fit either its x or y domain, but not both. 
    Row int64 `json:"row,omitempty"`
    
    // X 
    // info_array 
    // Sets the horizontal domain of this geo subplot (in plot fraction). Note that geo subplots are constrained by domain. In general, when `projection.scale` is set to 1. a map will fit either its x or y domain, but not both. 
    X interface{} `json:"x,omitempty"`
    
    // Y 
    // info_array 
    // Sets the vertical domain of this geo subplot (in plot fraction). Note that geo subplots are constrained by domain. In general, when `projection.scale` is set to 1. a map will fit either its x or y domain, but not both. 
    Y interface{} `json:"y,omitempty"`
    
}
// LayoutGeoLataxis 
type LayoutGeoLataxis struct {
    
    // Dtick 
    // number 
    // Sets the graticule's longitude/latitude tick step. 
    Dtick float64 `json:"dtick,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the graticule's stroke color. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the graticule's stroke width (in px). 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Range 
    // info_array 
    // Sets the range of this axis (in degrees), sets the map's clipped coordinates. 
    Range interface{} `json:"range,omitempty"`
    
    // Showgrid 
    // boolean 
    // Sets whether or not graticule are shown on the map. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Tick0 
    // number 
    // Sets the graticule's starting tick longitude/latitude. 
    Tick0 float64 `json:"tick0,omitempty"`
    
}
// LayoutGeoLonaxis 
type LayoutGeoLonaxis struct {
    
    // Dtick 
    // number 
    // Sets the graticule's longitude/latitude tick step. 
    Dtick float64 `json:"dtick,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the graticule's stroke color. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the graticule's stroke width (in px). 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Range 
    // info_array 
    // Sets the range of this axis (in degrees), sets the map's clipped coordinates. 
    Range interface{} `json:"range,omitempty"`
    
    // Showgrid 
    // boolean 
    // Sets whether or not graticule are shown on the map. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Tick0 
    // number 
    // Sets the graticule's starting tick longitude/latitude. 
    Tick0 float64 `json:"tick0,omitempty"`
    
}
// LayoutGeoProjectionRotation 
type LayoutGeoProjectionRotation struct {
    
    // Lat 
    // number 
    // Rotates the map along meridians (in degrees North). 
    Lat float64 `json:"lat,omitempty"`
    
    // Lon 
    // number 
    // Rotates the map along parallels (in degrees East). Defaults to the center of the `lonaxis.range` values. 
    Lon float64 `json:"lon,omitempty"`
    
    // Roll 
    // number 
    // Roll the map (in degrees) For example, a roll of *180* makes the map appear upside down. 
    Roll float64 `json:"roll,omitempty"`
    
}
// LayoutGeoProjection 
type LayoutGeoProjection struct {
    
    // Parallels 
    // info_array 
    // For conic projection types only. Sets the parallels (tangent, secant) where the cone intersects the sphere. 
    Parallels interface{} `json:"parallels,omitempty"`
    
    // Rotation 
    //   
    Rotation *LayoutGeoProjectionRotation `json:"rotation,omitempty"`
    
    // Scale 
    // number 
    // Zooms in or out on the map view. A scale of *1* corresponds to the largest zoom level that fits the map's lon and lat ranges.  
    Scale float64 `json:"scale,omitempty"`
    
    // Type 
    // enumerated Sets the projection type. 
    Type LayoutGeoProjectionType `json:"type,omitempty"`
    
}
// LayoutGeo 
type LayoutGeo struct {
    
    // Bgcolor 
    // color 
    // Set the background color of the map 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Center 
    //   
    Center *LayoutGeoCenter `json:"center,omitempty"`
    
    // Coastlinecolor 
    // color 
    // Sets the coastline color. 
    Coastlinecolor Color `json:"coastlinecolor,omitempty"`
    
    // Coastlinewidth 
    // number 
    // Sets the coastline stroke width (in px). 
    Coastlinewidth float64 `json:"coastlinewidth,omitempty"`
    
    // Countrycolor 
    // color 
    // Sets line color of the country boundaries. 
    Countrycolor Color `json:"countrycolor,omitempty"`
    
    // Countrywidth 
    // number 
    // Sets line width (in px) of the country boundaries. 
    Countrywidth float64 `json:"countrywidth,omitempty"`
    
    // Domain 
    //   
    Domain *LayoutGeoDomain `json:"domain,omitempty"`
    
    // Fitbounds 
    // enumerated Determines if this subplot's view settings are auto-computed to fit trace data. On scoped maps, setting `fitbounds` leads to `center.lon` and `center.lat` getting auto-filled. On maps with a non-clipped projection, setting `fitbounds` leads to `center.lon`, `center.lat`, and `projection.rotation.lon` getting auto-filled. On maps with a clipped projection, setting `fitbounds` leads to `center.lon`, `center.lat`, `projection.rotation.lon`, `projection.rotation.lat`, `lonaxis.range` and `lonaxis.range` getting auto-filled. If *locations*, only the trace's visible locations are considered in the `fitbounds` computations. If *geojson*, the entire trace input `geojson` (if provided) is considered in the `fitbounds` computations, Defaults to *false*. 
    Fitbounds LayoutGeoFitbounds `json:"fitbounds,omitempty"`
    
    // Framecolor 
    // color 
    // Sets the color the frame. 
    Framecolor Color `json:"framecolor,omitempty"`
    
    // Framewidth 
    // number 
    // Sets the stroke width (in px) of the frame. 
    Framewidth float64 `json:"framewidth,omitempty"`
    
    // Lakecolor 
    // color 
    // Sets the color of the lakes. 
    Lakecolor Color `json:"lakecolor,omitempty"`
    
    // Landcolor 
    // color 
    // Sets the land mass color. 
    Landcolor Color `json:"landcolor,omitempty"`
    
    // Lataxis 
    //   
    Lataxis *LayoutGeoLataxis `json:"lataxis,omitempty"`
    
    // Lonaxis 
    //   
    Lonaxis *LayoutGeoLonaxis `json:"lonaxis,omitempty"`
    
    // Oceancolor 
    // color 
    // Sets the ocean color 
    Oceancolor Color `json:"oceancolor,omitempty"`
    
    // Projection 
    //   
    Projection *LayoutGeoProjection `json:"projection,omitempty"`
    
    // Resolution 
    // enumerated Sets the resolution of the base layers. The values have units of km/mm e.g. 110 corresponds to a scale ratio of 1:110,000,000. 
    Resolution LayoutGeoResolution `json:"resolution,omitempty"`
    
    // Rivercolor 
    // color 
    // Sets color of the rivers. 
    Rivercolor Color `json:"rivercolor,omitempty"`
    
    // Riverwidth 
    // number 
    // Sets the stroke width (in px) of the rivers. 
    Riverwidth float64 `json:"riverwidth,omitempty"`
    
    // Scope 
    // enumerated Set the scope of the map. 
    Scope LayoutGeoScope `json:"scope,omitempty"`
    
    // Showcoastlines 
    // boolean 
    // Sets whether or not the coastlines are drawn. 
    Showcoastlines Bool `json:"showcoastlines,omitempty"`
    
    // Showcountries 
    // boolean 
    // Sets whether or not country boundaries are drawn. 
    Showcountries Bool `json:"showcountries,omitempty"`
    
    // Showframe 
    // boolean 
    // Sets whether or not a frame is drawn around the map. 
    Showframe Bool `json:"showframe,omitempty"`
    
    // Showlakes 
    // boolean 
    // Sets whether or not lakes are drawn. 
    Showlakes Bool `json:"showlakes,omitempty"`
    
    // Showland 
    // boolean 
    // Sets whether or not land masses are filled in color. 
    Showland Bool `json:"showland,omitempty"`
    
    // Showocean 
    // boolean 
    // Sets whether or not oceans are filled in color. 
    Showocean Bool `json:"showocean,omitempty"`
    
    // Showrivers 
    // boolean 
    // Sets whether or not rivers are drawn. 
    Showrivers Bool `json:"showrivers,omitempty"`
    
    // Showsubunits 
    // boolean 
    // Sets whether or not boundaries of subunits within countries (e.g. states, provinces) are drawn. 
    Showsubunits Bool `json:"showsubunits,omitempty"`
    
    // Subunitcolor 
    // color 
    // Sets the color of the subunits boundaries. 
    Subunitcolor Color `json:"subunitcolor,omitempty"`
    
    // Subunitwidth 
    // number 
    // Sets the stroke width (in px) of the subunits boundaries. 
    Subunitwidth float64 `json:"subunitwidth,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in the view (projection and center). Defaults to `layout.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Visible 
    // boolean 
    // Sets the default visibility of the base layers. 
    Visible Bool `json:"visible,omitempty"`
    
}
// LayoutGridDomain 
type LayoutGridDomain struct {
    
    // X 
    // info_array 
    // Sets the horizontal domain of this grid subplot (in plot fraction). The first and last cells end exactly at the domain edges, with no grout around the edges. 
    X interface{} `json:"x,omitempty"`
    
    // Y 
    // info_array 
    // Sets the vertical domain of this grid subplot (in plot fraction). The first and last cells end exactly at the domain edges, with no grout around the edges. 
    Y interface{} `json:"y,omitempty"`
    
}
// LayoutGrid 
type LayoutGrid struct {
    
    // Columns 
    // integer 
    // The number of columns in the grid. If you provide a 2D `subplots` array, the length of its longest row is used as the default. If you give an `xaxes` array, its length is used as the default. But it's also possible to have a different length, if you want to leave a row at the end for non-cartesian subplots. 
    Columns int64 `json:"columns,omitempty"`
    
    // Domain 
    //   
    Domain *LayoutGridDomain `json:"domain,omitempty"`
    
    // Pattern 
    // enumerated If no `subplots`, `xaxes`, or `yaxes` are given but we do have `rows` and `columns`, we can generate defaults using consecutive axis IDs, in two ways: *coupled* gives one x axis per column and one y axis per row. *independent* uses a new xy pair for each cell, left-to-right across each row then iterating rows according to `roworder`. 
    Pattern LayoutGridPattern `json:"pattern,omitempty"`
    
    // Roworder 
    // enumerated Is the first row the top or the bottom? Note that columns are always enumerated from left to right. 
    Roworder LayoutGridRoworder `json:"roworder,omitempty"`
    
    // Rows 
    // integer 
    // The number of rows in the grid. If you provide a 2D `subplots` array or a `yaxes` array, its length is used as the default. But it's also possible to have a different length, if you want to leave a row at the end for non-cartesian subplots. 
    Rows int64 `json:"rows,omitempty"`
    
    // Subplots 
    // info_array 
    // Used for freeform grids, where some axes may be shared across subplots but others are not. Each entry should be a cartesian subplot id, like *xy* or *x3y2*, or ** to leave that cell empty. You may reuse x axes within the same column, and y axes within the same row. Non-cartesian subplots and traces that support `domain` can place themselves in this grid separately using the `gridcell` attribute. 
    Subplots interface{} `json:"subplots,omitempty"`
    
    // Xaxes 
    // info_array 
    // Used with `yaxes` when the x and y axes are shared across columns and rows. Each entry should be an x axis id like *x*, *x2*, etc., or ** to not put an x axis in that column. Entries other than ** must be unique. Ignored if `subplots` is present. If missing but `yaxes` is present, will generate consecutive IDs. 
    Xaxes interface{} `json:"xaxes,omitempty"`
    
    // Xgap 
    // number 
    // Horizontal space between grid cells, expressed as a fraction of the total width available to one cell. Defaults to 0.1 for coupled-axes grids and 0.2 for independent grids. 
    Xgap float64 `json:"xgap,omitempty"`
    
    // Xside 
    // enumerated Sets where the x axis labels and titles go. *bottom* means the very bottom of the grid. *bottom plot* is the lowest plot that each x axis is used in. *top* and *top plot* are similar. 
    Xside LayoutGridXside `json:"xside,omitempty"`
    
    // Yaxes 
    // info_array 
    // Used with `yaxes` when the x and y axes are shared across columns and rows. Each entry should be an y axis id like *y*, *y2*, etc., or ** to not put a y axis in that row. Entries other than ** must be unique. Ignored if `subplots` is present. If missing but `xaxes` is present, will generate consecutive IDs. 
    Yaxes interface{} `json:"yaxes,omitempty"`
    
    // Ygap 
    // number 
    // Vertical space between grid cells, expressed as a fraction of the total height available to one cell. Defaults to 0.1 for coupled-axes grids and 0.3 for independent grids. 
    Ygap float64 `json:"ygap,omitempty"`
    
    // Yside 
    // enumerated Sets where the y axis labels and titles go. *left* means the very left edge of the grid. *left plot* is the leftmost plot that each y axis is used in. *right* and *right plot* are similar. 
    Yside LayoutGridYside `json:"yside,omitempty"`
    
}
// LayoutHoverlabelFont Sets the default hover label font used by all traces on the graph.
type LayoutHoverlabelFont struct {
    
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
// LayoutHoverlabel 
type LayoutHoverlabel struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines 
    Align LayoutHoverlabelAlign `json:"align,omitempty"`
    
    // Bgcolor 
    // color 
    // Sets the background color of all hover labels on graph 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Bordercolor 
    // color 
    // Sets the border color of all hover labels on graph. 
    Bordercolor Color `json:"bordercolor,omitempty"`
    
    // Font 
    //  Sets the default hover label font used by all traces on the graph. 
    Font *LayoutHoverlabelFont `json:"font,omitempty"`
    
    // Namelength 
    // integer 
    // Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis. 
    Namelength int64 `json:"namelength,omitempty"`
    
}
// LayoutLegendFont Sets the font used to text the legend items.
type LayoutLegendFont struct {
    
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
// LayoutLegendTitleFont Sets this legend's title font.
type LayoutLegendTitleFont struct {
    
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
// LayoutLegendTitle 
type LayoutLegendTitle struct {
    
    // Font 
    //  Sets this legend's title font. 
    Font *LayoutLegendTitleFont `json:"font,omitempty"`
    
    // Side 
    // enumerated Determines the location of legend's title with respect to the legend items. Defaulted to *top* with `orientation` is *h*. Defaulted to *left* with `orientation` is *v*. The *top left* options could be used to expand legend area in both x and y sides. 
    Side LayoutLegendTitleSide `json:"side,omitempty"`
    
    // Text 
    // string 
    // Sets the title of the legend. 
    Text String `json:"text,omitempty"`
    
}
// LayoutLegend 
type LayoutLegend struct {
    
    // Bgcolor 
    // color 
    // Sets the legend background color. Defaults to `layout.paper_bgcolor`. 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Bordercolor 
    // color 
    // Sets the color of the border enclosing the legend. 
    Bordercolor Color `json:"bordercolor,omitempty"`
    
    // Borderwidth 
    // number 
    // Sets the width (in px) of the border enclosing the legend. 
    Borderwidth float64 `json:"borderwidth,omitempty"`
    
    // Font 
    //  Sets the font used to text the legend items. 
    Font *LayoutLegendFont `json:"font,omitempty"`
    
    // Itemclick 
    // enumerated Determines the behavior on legend item click. *toggle* toggles the visibility of the item clicked on the graph. *toggleothers* makes the clicked item the sole visible item on the graph. *false* disable legend item click interactions. 
    Itemclick LayoutLegendItemclick `json:"itemclick,omitempty"`
    
    // Itemdoubleclick 
    // enumerated Determines the behavior on legend item double-click. *toggle* toggles the visibility of the item clicked on the graph. *toggleothers* makes the clicked item the sole visible item on the graph. *false* disable legend item double-click interactions. 
    Itemdoubleclick LayoutLegendItemdoubleclick `json:"itemdoubleclick,omitempty"`
    
    // Itemsizing 
    // enumerated Determines if the legend items symbols scale with their corresponding *trace* attributes or remain *constant* independent of the symbol size on the graph. 
    Itemsizing LayoutLegendItemsizing `json:"itemsizing,omitempty"`
    
    // Itemwidth 
    // number 
    // Sets the width (in px) of the legend item symbols (the part other than the title.text). 
    Itemwidth float64 `json:"itemwidth,omitempty"`
    
    // Orientation 
    // enumerated Sets the orientation of the legend. 
    Orientation LayoutLegendOrientation `json:"orientation,omitempty"`
    
    // Title 
    //   
    Title *LayoutLegendTitle `json:"title,omitempty"`
    
    // Tracegroupgap 
    // number 
    // Sets the amount of vertical space (in px) between legend groups. 
    Tracegroupgap float64 `json:"tracegroupgap,omitempty"`
    
    // Traceorder 
    // flaglist Determines the order at which the legend items are displayed. If *normal*, the items are displayed top-to-bottom in the same order as the input data. If *reversed*, the items are displayed in the opposite order as *normal*. If *grouped*, the items are displayed in groups (when a trace `legendgroup` is provided). if *grouped+reversed*, the items are displayed in the opposite order as *grouped*. 
    Traceorder LayoutLegendTraceorder `json:"traceorder,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of legend-driven changes in trace and pie label visibility. Defaults to `layout.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Valign 
    // enumerated Sets the vertical alignment of the symbols with respect to their associated text. 
    Valign LayoutLegendValign `json:"valign,omitempty"`
    
    // X 
    // number 
    // Sets the x position (in normalized coordinates) of the legend. Defaults to *1.02* for vertical legends and defaults to *0* for horizontal legends. 
    X float64 `json:"x,omitempty"`
    
    // Xanchor 
    // enumerated Sets the legend's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the legend. Value *auto* anchors legends to the right for `x` values greater than or equal to 2/3, anchors legends to the left for `x` values less than or equal to 1/3 and anchors legends with respect to their center otherwise. 
    Xanchor LayoutLegendXanchor `json:"xanchor,omitempty"`
    
    // Y 
    // number 
    // Sets the y position (in normalized coordinates) of the legend. Defaults to *1* for vertical legends, defaults to *-0.1* for horizontal legends on graphs w/o range sliders and defaults to *1.1* for horizontal legends on graph with one or multiple range sliders. 
    Y float64 `json:"y,omitempty"`
    
    // Yanchor 
    // enumerated Sets the legend's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the legend. Value *auto* anchors legends at their bottom for `y` values less than or equal to 1/3, anchors legends to at their top for `y` values greater than or equal to 2/3 and anchors legends with respect to their middle otherwise. 
    Yanchor LayoutLegendYanchor `json:"yanchor,omitempty"`
    
}
// LayoutMapboxCenter 
type LayoutMapboxCenter struct {
    
    // Lat 
    // number 
    // Sets the latitude of the center of the map (in degrees North). 
    Lat float64 `json:"lat,omitempty"`
    
    // Lon 
    // number 
    // Sets the longitude of the center of the map (in degrees East). 
    Lon float64 `json:"lon,omitempty"`
    
}
// LayoutMapboxDomain 
type LayoutMapboxDomain struct {
    
    // Column 
    // integer 
    // If there is a layout grid, use the domain for this column in the grid for this mapbox subplot . 
    Column int64 `json:"column,omitempty"`
    
    // Row 
    // integer 
    // If there is a layout grid, use the domain for this row in the grid for this mapbox subplot . 
    Row int64 `json:"row,omitempty"`
    
    // X 
    // info_array 
    // Sets the horizontal domain of this mapbox subplot (in plot fraction). 
    X interface{} `json:"x,omitempty"`
    
    // Y 
    // info_array 
    // Sets the vertical domain of this mapbox subplot (in plot fraction). 
    Y interface{} `json:"y,omitempty"`
    
}
// LayoutMapbox 
type LayoutMapbox struct {
    
    // Accesstoken 
    // string 
    // Sets the mapbox access token to be used for this mapbox map. Alternatively, the mapbox access token can be set in the configuration options under `mapboxAccessToken`. Note that accessToken are only required when `style` (e.g with values : basic, streets, outdoors, light, dark, satellite, satellite-streets ) and/or a layout layer references the Mapbox server. 
    Accesstoken String `json:"accesstoken,omitempty"`
    
    // Bearing 
    // number 
    // Sets the bearing angle of the map in degrees counter-clockwise from North (mapbox.bearing). 
    Bearing float64 `json:"bearing,omitempty"`
    
    // Center 
    //   
    Center *LayoutMapboxCenter `json:"center,omitempty"`
    
    // Domain 
    //   
    Domain *LayoutMapboxDomain `json:"domain,omitempty"`
    
    // Layers 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Layers interface{} `json:"layers,omitempty"`
    
    // Pitch 
    // number 
    // Sets the pitch angle of the map (in degrees, where *0* means perpendicular to the surface of the map) (mapbox.pitch). 
    Pitch float64 `json:"pitch,omitempty"`
    
    // Style 
    // any 
    // Defines the map layers that are rendered by default below the trace layers defined in `data`, which are themselves by default rendered below the layers defined in `layout.mapbox.layers`.  These layers can be defined either explicitly as a Mapbox Style object which can contain multiple layer definitions that load data from any public or private Tile Map Service (TMS or XYZ) or Web Map Service (WMS) or implicitly by using one of the built-in style objects which use WMSes which do not require any access tokens, or by using a default Mapbox style or custom Mapbox style URL, both of which require a Mapbox access token  Note that Mapbox access token can be set in the `accesstoken` attribute or in the `mapboxAccessToken` config option.  Mapbox Style objects are of the form described in the Mapbox GL JS documentation available at https://docs.mapbox.com/mapbox-gl-js/style-spec  The built-in plotly.js styles objects are: open-street-map, white-bg, carto-positron, carto-darkmatter, stamen-terrain, stamen-toner, stamen-watercolor  The built-in Mapbox styles are: basic, streets, outdoors, light, dark, satellite, satellite-streets  Mapbox style URLs are of the form: mapbox://mapbox.mapbox-<name>-<version> 
    Style interface{} `json:"style,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in the view: `center`, `zoom`, `bearing`, `pitch`. Defaults to `layout.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Zoom 
    // number 
    // Sets the zoom level of the map (mapbox.zoom). 
    Zoom float64 `json:"zoom,omitempty"`
    
}
// LayoutMargin 
type LayoutMargin struct {
    
    // Autoexpand 
    // boolean 
    // Turns on/off margin expansion computations. Legends, colorbars, updatemenus, sliders, axis rangeselector and rangeslider are allowed to push the margins by defaults. 
    Autoexpand Bool `json:"autoexpand,omitempty"`
    
    // B 
    // number 
    // Sets the bottom margin (in px). 
    B float64 `json:"b,omitempty"`
    
    // L 
    // number 
    // Sets the left margin (in px). 
    L float64 `json:"l,omitempty"`
    
    // Pad 
    // number 
    // Sets the amount of padding (in px) between the plotting area and the axis lines 
    Pad float64 `json:"pad,omitempty"`
    
    // R 
    // number 
    // Sets the right margin (in px). 
    R float64 `json:"r,omitempty"`
    
    // T 
    // number 
    // Sets the top margin (in px). 
    T float64 `json:"t,omitempty"`
    
}
// LayoutModebar 
type LayoutModebar struct {
    
    // Activecolor 
    // color 
    // Sets the color of the active or hovered on icons in the modebar. 
    Activecolor Color `json:"activecolor,omitempty"`
    
    // Bgcolor 
    // color 
    // Sets the background color of the modebar. 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Color 
    // color 
    // Sets the color of the icons in the modebar. 
    Color Color `json:"color,omitempty"`
    
    // Orientation 
    // enumerated Sets the orientation of the modebar. 
    Orientation LayoutModebarOrientation `json:"orientation,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes related to the modebar, including `hovermode`, `dragmode`, and `showspikes` at both the root level and inside subplots. Defaults to `layout.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
}
// LayoutNewshapeLine 
type LayoutNewshapeLine struct {
    
    // Color 
    // color 
    // Sets the line color. By default uses either dark grey or white to increase contrast with background color. 
    Color Color `json:"color,omitempty"`
    
    // Dash 
    // string 
    // Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*). 
    Dash String `json:"dash,omitempty"`
    
    // Width 
    // number 
    // Sets the line width (in px). 
    Width float64 `json:"width,omitempty"`
    
}
// LayoutNewshape 
type LayoutNewshape struct {
    
    // Drawdirection 
    // enumerated When `dragmode` is set to *drawrect*, *drawline* or *drawcircle* this limits the drag to be horizontal, vertical or diagonal. Using *diagonal* there is no limit e.g. in drawing lines in any direction. *ortho* limits the draw to be either horizontal or vertical. *horizontal* allows horizontal extend. *vertical* allows vertical extend. 
    Drawdirection LayoutNewshapeDrawdirection `json:"drawdirection,omitempty"`
    
    // Fillcolor 
    // color 
    // Sets the color filling new shapes' interior. Please note that if using a fillcolor with alpha greater than half, drag inside the active shape starts moving the shape underneath, otherwise a new shape could be started over. 
    Fillcolor Color `json:"fillcolor,omitempty"`
    
    // Fillrule 
    // enumerated Determines the path's interior. For more info please visit https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill-rule 
    Fillrule LayoutNewshapeFillrule `json:"fillrule,omitempty"`
    
    // Layer 
    // enumerated Specifies whether new shapes are drawn below or above traces. 
    Layer LayoutNewshapeLayer `json:"layer,omitempty"`
    
    // Line 
    //   
    Line *LayoutNewshapeLine `json:"line,omitempty"`
    
    // Opacity 
    // number 
    // Sets the opacity of new shapes. 
    Opacity float64 `json:"opacity,omitempty"`
    
}
// LayoutPolarAngularaxisTickfont Sets the tick font.
type LayoutPolarAngularaxisTickfont struct {
    
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
// LayoutPolarAngularaxis 
type LayoutPolarAngularaxis struct {
    
    // Autotypenumbers 
    // enumerated Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers. 
    Autotypenumbers LayoutPolarAngularaxisAutotypenumbers `json:"autotypenumbers,omitempty"`
    
    // Categoryarray 
    // data_array 
    // Sets the order in which categories on this axis appear. Only has an effect if `categoryorder` is set to *array*. Used with `categoryorder`. 
    Categoryarray interface{} `json:"categoryarray,omitempty"`
    
    // Categoryarraysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  categoryarray . 
    Categoryarraysrc String `json:"categoryarraysrc,omitempty"`
    
    // Categoryorder 
    // enumerated Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values. 
    Categoryorder LayoutPolarAngularaxisCategoryorder `json:"categoryorder,omitempty"`
    
    // Color 
    // color 
    // Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this. 
    Color Color `json:"color,omitempty"`
    
    // Direction 
    // enumerated Sets the direction corresponding to positive angles. 
    Direction LayoutPolarAngularaxisDirection `json:"direction,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat LayoutPolarAngularaxisExponentformat `json:"exponentformat,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the color of the grid lines. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the width (in px) of the grid lines. 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Hoverformat 
    // string 
    // Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Hoverformat String `json:"hoverformat,omitempty"`
    
    // Layer 
    // enumerated Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis. 
    Layer LayoutPolarAngularaxisLayer `json:"layer,omitempty"`
    
    // Linecolor 
    // color 
    // Sets the axis line color. 
    Linecolor Color `json:"linecolor,omitempty"`
    
    // Linewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Linewidth float64 `json:"linewidth,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Period 
    // number 
    // Set the angular period. Has an effect only when `angularaxis.type` is *category*. 
    Period float64 `json:"period,omitempty"`
    
    // Rotation 
    // angle 
    // Sets that start position (in degrees) of the angular axis By default, polar subplots with `direction` set to *counterclockwise* get a `rotation` of *0* which corresponds to due East (like what mathematicians prefer). In turn, polar with `direction` set to *clockwise* get a rotation of *90* which corresponds to due North (like on a compass), 
    Rotation float64 `json:"rotation,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent LayoutPolarAngularaxisShowexponent `json:"showexponent,omitempty"`
    
    // Showgrid 
    // boolean 
    // Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Showline 
    // boolean 
    // Determines whether or not a line bounding this axis is drawn. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutPolarAngularaxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutPolarAngularaxisShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Thetaunit 
    // enumerated Sets the format unit of the formatted *theta* values. Has an effect only when `angularaxis.type` is *linear*. 
    Thetaunit LayoutPolarAngularaxisThetaunit `json:"thetaunit,omitempty"`
    
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
    //  Sets the tick font. 
    Tickfont *LayoutPolarAngularaxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutPolarAngularaxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutPolarAngularaxisTicks `json:"ticks,omitempty"`
    
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
    
    // Type 
    // enumerated Sets the angular axis type. If *linear*, set `thetaunit` to determine the unit in which axis value are shown. If *category, use `period` to set the number of integer coordinates around polar axis. 
    Type LayoutPolarAngularaxisType `json:"type,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in axis `rotation`. Defaults to `polar<N>.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Visible 
    // boolean 
    // A single toggle to hide the axis while preserving interaction like dragging. Default is true when a cheater plot is present on the axis, otherwise false 
    Visible Bool `json:"visible,omitempty"`
    
}
// LayoutPolarDomain 
type LayoutPolarDomain struct {
    
    // Column 
    // integer 
    // If there is a layout grid, use the domain for this column in the grid for this polar subplot . 
    Column int64 `json:"column,omitempty"`
    
    // Row 
    // integer 
    // If there is a layout grid, use the domain for this row in the grid for this polar subplot . 
    Row int64 `json:"row,omitempty"`
    
    // X 
    // info_array 
    // Sets the horizontal domain of this polar subplot (in plot fraction). 
    X interface{} `json:"x,omitempty"`
    
    // Y 
    // info_array 
    // Sets the vertical domain of this polar subplot (in plot fraction). 
    Y interface{} `json:"y,omitempty"`
    
}
// LayoutPolarRadialaxisTickfont Sets the tick font.
type LayoutPolarRadialaxisTickfont struct {
    
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
// LayoutPolarRadialaxisTitleFont Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute.
type LayoutPolarRadialaxisTitleFont struct {
    
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
// LayoutPolarRadialaxisTitle 
type LayoutPolarRadialaxisTitle struct {
    
    // Font 
    //  Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute. 
    Font *LayoutPolarRadialaxisTitleFont `json:"font,omitempty"`
    
    // Text 
    // string 
    // Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// LayoutPolarRadialaxis 
type LayoutPolarRadialaxis struct {
    
    // Angle 
    // angle 
    // Sets the angle (in degrees) from which the radial axis is drawn. Note that by default, radial axis line on the theta=0 line corresponds to a line pointing right (like what mathematicians prefer). Defaults to the first `polar.sector` angle. 
    Angle float64 `json:"angle,omitempty"`
    
    // Autorange 
    // enumerated Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*. 
    Autorange LayoutPolarRadialaxisAutorange `json:"autorange,omitempty"`
    
    // Autotypenumbers 
    // enumerated Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers. 
    Autotypenumbers LayoutPolarRadialaxisAutotypenumbers `json:"autotypenumbers,omitempty"`
    
    // Calendar 
    // enumerated Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar` 
    Calendar LayoutPolarRadialaxisCalendar `json:"calendar,omitempty"`
    
    // Categoryarray 
    // data_array 
    // Sets the order in which categories on this axis appear. Only has an effect if `categoryorder` is set to *array*. Used with `categoryorder`. 
    Categoryarray interface{} `json:"categoryarray,omitempty"`
    
    // Categoryarraysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  categoryarray . 
    Categoryarraysrc String `json:"categoryarraysrc,omitempty"`
    
    // Categoryorder 
    // enumerated Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values. 
    Categoryorder LayoutPolarRadialaxisCategoryorder `json:"categoryorder,omitempty"`
    
    // Color 
    // color 
    // Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this. 
    Color Color `json:"color,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat LayoutPolarRadialaxisExponentformat `json:"exponentformat,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the color of the grid lines. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the width (in px) of the grid lines. 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Hoverformat 
    // string 
    // Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Hoverformat String `json:"hoverformat,omitempty"`
    
    // Layer 
    // enumerated Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis. 
    Layer LayoutPolarRadialaxisLayer `json:"layer,omitempty"`
    
    // Linecolor 
    // color 
    // Sets the axis line color. 
    Linecolor Color `json:"linecolor,omitempty"`
    
    // Linewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Linewidth float64 `json:"linewidth,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Range 
    // info_array 
    // Sets the range of this axis. If the axis `type` is *log*, then you must take the log of your desired range (e.g. to set the range from 1 to 100, set the range from 0 to 2). If the axis `type` is *date*, it should be date strings, like date data, though Date objects and unix milliseconds will be accepted and converted to strings. If the axis `type` is *category*, it should be numbers, using the scale where each category is assigned a serial number from zero in the order it appears. 
    Range interface{} `json:"range,omitempty"`
    
    // Rangemode 
    // enumerated If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. If *normal*, the range is computed in relation to the extrema of the input data (same behavior as for cartesian axes). 
    Rangemode LayoutPolarRadialaxisRangemode `json:"rangemode,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent LayoutPolarRadialaxisShowexponent `json:"showexponent,omitempty"`
    
    // Showgrid 
    // boolean 
    // Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Showline 
    // boolean 
    // Determines whether or not a line bounding this axis is drawn. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutPolarRadialaxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutPolarRadialaxisShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Side 
    // enumerated Determines on which side of radial axis line the tick and tick labels appear. 
    Side LayoutPolarRadialaxisSide `json:"side,omitempty"`
    
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
    //  Sets the tick font. 
    Tickfont *LayoutPolarRadialaxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutPolarRadialaxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutPolarRadialaxisTicks `json:"ticks,omitempty"`
    
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
    Title *LayoutPolarRadialaxisTitle `json:"title,omitempty"`
    
    // Type 
    // enumerated Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question. 
    Type LayoutPolarRadialaxisType `json:"type,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in axis `range`, `autorange`, `angle`, and `title` if in `editable: true` configuration. Defaults to `polar<N>.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Visible 
    // boolean 
    // A single toggle to hide the axis while preserving interaction like dragging. Default is true when a cheater plot is present on the axis, otherwise false 
    Visible Bool `json:"visible,omitempty"`
    
}
// LayoutPolar 
type LayoutPolar struct {
    
    // Angularaxis 
    //   
    Angularaxis *LayoutPolarAngularaxis `json:"angularaxis,omitempty"`
    
    // Bgcolor 
    // color 
    // Set the background color of the subplot 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Domain 
    //   
    Domain *LayoutPolarDomain `json:"domain,omitempty"`
    
    // Gridshape 
    // enumerated Determines if the radial axis grid lines and angular axis line are drawn as *circular* sectors or as *linear* (polygon) sectors. Has an effect only when the angular axis has `type` *category*. Note that `radialaxis.angle` is snapped to the angle of the closest vertex when `gridshape` is *circular* (so that radial axis scale is the same as the data scale). 
    Gridshape LayoutPolarGridshape `json:"gridshape,omitempty"`
    
    // Hole 
    // number 
    // Sets the fraction of the radius to cut out of the polar subplot. 
    Hole float64 `json:"hole,omitempty"`
    
    // Radialaxis 
    //   
    Radialaxis *LayoutPolarRadialaxis `json:"radialaxis,omitempty"`
    
    // Sector 
    // info_array 
    // Sets angular span of this polar subplot with two angles (in degrees). Sector are assumed to be spanned in the counterclockwise direction with *0* corresponding to rightmost limit of the polar subplot. 
    Sector interface{} `json:"sector,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in axis attributes, if not overridden in the individual axes. Defaults to `layout.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
}
// LayoutRadialaxis 
type LayoutRadialaxis struct {
    
    // Domain 
    // info_array 
    // Polar chart subplots are not supported yet. This key has currently no effect. 
    Domain interface{} `json:"domain,omitempty"`
    
    // Endpadding 
    // number 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. 
    Endpadding float64 `json:"endpadding,omitempty"`
    
    // Orientation 
    // number 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the orientation (an angle with respect to the origin) of the radial axis. 
    Orientation float64 `json:"orientation,omitempty"`
    
    // Range 
    // info_array 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Defines the start and end point of this radial axis. 
    Range interface{} `json:"range,omitempty"`
    
    // Showline 
    // boolean 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Determines whether or not the line bounding this radial axis will be shown on the figure. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Determines whether or not the radial axis ticks will feature tick labels. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Tickcolor 
    // color 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the color of the tick lines on this radial axis. 
    Tickcolor Color `json:"tickcolor,omitempty"`
    
    // Ticklen 
    // number 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the length of the tick lines on this radial axis. 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickorientation 
    // enumerated Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the orientation (from the paper perspective) of the radial axis tick labels. 
    Tickorientation LayoutRadialaxisTickorientation `json:"tickorientation,omitempty"`
    
    // Ticksuffix 
    // string 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the length of the tick lines on this radial axis. 
    Ticksuffix String `json:"ticksuffix,omitempty"`
    
    // Visible 
    // boolean 
    // Legacy polar charts are deprecated! Please switch to *polar* subplots. Determines whether or not this axis will be visible. 
    Visible Bool `json:"visible,omitempty"`
    
}
// LayoutSceneAspectratio Sets this scene's axis aspectratio.
type LayoutSceneAspectratio struct {
    
    // X 
    // number 
    //  
    X float64 `json:"x,omitempty"`
    
    // Y 
    // number 
    //  
    Y float64 `json:"y,omitempty"`
    
    // Z 
    // number 
    //  
    Z float64 `json:"z,omitempty"`
    
}
// LayoutSceneCameraCenter Sets the (x,y,z) components of the 'center' camera vector This vector determines the translation (x,y,z) space about the center of this scene. By default, there is no such translation.
type LayoutSceneCameraCenter struct {
    
    // X 
    // number 
    //  
    X float64 `json:"x,omitempty"`
    
    // Y 
    // number 
    //  
    Y float64 `json:"y,omitempty"`
    
    // Z 
    // number 
    //  
    Z float64 `json:"z,omitempty"`
    
}
// LayoutSceneCameraEye Sets the (x,y,z) components of the 'eye' camera vector. This vector determines the view point about the origin of this scene.
type LayoutSceneCameraEye struct {
    
    // X 
    // number 
    //  
    X float64 `json:"x,omitempty"`
    
    // Y 
    // number 
    //  
    Y float64 `json:"y,omitempty"`
    
    // Z 
    // number 
    //  
    Z float64 `json:"z,omitempty"`
    
}
// LayoutSceneCameraProjection 
type LayoutSceneCameraProjection struct {
    
    // Type 
    // enumerated Sets the projection type. The projection type could be either *perspective* or *orthographic*. The default is *perspective*. 
    Type LayoutSceneCameraProjectionType `json:"type,omitempty"`
    
}
// LayoutSceneCameraUp Sets the (x,y,z) components of the 'up' camera vector. This vector determines the up direction of this scene with respect to the page. The default is *{x: 0, y: 0, z: 1}* which means that the z axis points up.
type LayoutSceneCameraUp struct {
    
    // X 
    // number 
    //  
    X float64 `json:"x,omitempty"`
    
    // Y 
    // number 
    //  
    Y float64 `json:"y,omitempty"`
    
    // Z 
    // number 
    //  
    Z float64 `json:"z,omitempty"`
    
}
// LayoutSceneCamera 
type LayoutSceneCamera struct {
    
    // Center 
    //  Sets the (x,y,z) components of the 'center' camera vector This vector determines the translation (x,y,z) space about the center of this scene. By default, there is no such translation. 
    Center *LayoutSceneCameraCenter `json:"center,omitempty"`
    
    // Eye 
    //  Sets the (x,y,z) components of the 'eye' camera vector. This vector determines the view point about the origin of this scene. 
    Eye *LayoutSceneCameraEye `json:"eye,omitempty"`
    
    // Projection 
    //   
    Projection *LayoutSceneCameraProjection `json:"projection,omitempty"`
    
    // Up 
    //  Sets the (x,y,z) components of the 'up' camera vector. This vector determines the up direction of this scene with respect to the page. The default is *{x: 0, y: 0, z: 1}* which means that the z axis points up. 
    Up *LayoutSceneCameraUp `json:"up,omitempty"`
    
}
// LayoutSceneDomain 
type LayoutSceneDomain struct {
    
    // Column 
    // integer 
    // If there is a layout grid, use the domain for this column in the grid for this scene subplot . 
    Column int64 `json:"column,omitempty"`
    
    // Row 
    // integer 
    // If there is a layout grid, use the domain for this row in the grid for this scene subplot . 
    Row int64 `json:"row,omitempty"`
    
    // X 
    // info_array 
    // Sets the horizontal domain of this scene subplot (in plot fraction). 
    X interface{} `json:"x,omitempty"`
    
    // Y 
    // info_array 
    // Sets the vertical domain of this scene subplot (in plot fraction). 
    Y interface{} `json:"y,omitempty"`
    
}
// LayoutSceneXaxisTickfont Sets the tick font.
type LayoutSceneXaxisTickfont struct {
    
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
// LayoutSceneXaxisTitleFont Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute.
type LayoutSceneXaxisTitleFont struct {
    
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
// LayoutSceneXaxisTitle 
type LayoutSceneXaxisTitle struct {
    
    // Font 
    //  Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute. 
    Font *LayoutSceneXaxisTitleFont `json:"font,omitempty"`
    
    // Text 
    // string 
    // Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// LayoutSceneXaxis 
type LayoutSceneXaxis struct {
    
    // Autorange 
    // enumerated Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*. 
    Autorange LayoutSceneXaxisAutorange `json:"autorange,omitempty"`
    
    // Autotypenumbers 
    // enumerated Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers. 
    Autotypenumbers LayoutSceneXaxisAutotypenumbers `json:"autotypenumbers,omitempty"`
    
    // Backgroundcolor 
    // color 
    // Sets the background color of this axis' wall. 
    Backgroundcolor Color `json:"backgroundcolor,omitempty"`
    
    // Calendar 
    // enumerated Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar` 
    Calendar LayoutSceneXaxisCalendar `json:"calendar,omitempty"`
    
    // Categoryarray 
    // data_array 
    // Sets the order in which categories on this axis appear. Only has an effect if `categoryorder` is set to *array*. Used with `categoryorder`. 
    Categoryarray interface{} `json:"categoryarray,omitempty"`
    
    // Categoryarraysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  categoryarray . 
    Categoryarraysrc String `json:"categoryarraysrc,omitempty"`
    
    // Categoryorder 
    // enumerated Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values. 
    Categoryorder LayoutSceneXaxisCategoryorder `json:"categoryorder,omitempty"`
    
    // Color 
    // color 
    // Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this. 
    Color Color `json:"color,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat LayoutSceneXaxisExponentformat `json:"exponentformat,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the color of the grid lines. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the width (in px) of the grid lines. 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Hoverformat 
    // string 
    // Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Hoverformat String `json:"hoverformat,omitempty"`
    
    // Linecolor 
    // color 
    // Sets the axis line color. 
    Linecolor Color `json:"linecolor,omitempty"`
    
    // Linewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Linewidth float64 `json:"linewidth,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Mirror 
    // enumerated Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots. 
    Mirror LayoutSceneXaxisMirror `json:"mirror,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Range 
    // info_array 
    // Sets the range of this axis. If the axis `type` is *log*, then you must take the log of your desired range (e.g. to set the range from 1 to 100, set the range from 0 to 2). If the axis `type` is *date*, it should be date strings, like date data, though Date objects and unix milliseconds will be accepted and converted to strings. If the axis `type` is *category*, it should be numbers, using the scale where each category is assigned a serial number from zero in the order it appears. 
    Range interface{} `json:"range,omitempty"`
    
    // Rangemode 
    // enumerated If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes. 
    Rangemode LayoutSceneXaxisRangemode `json:"rangemode,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showaxeslabels 
    // boolean 
    // Sets whether or not this axis is labeled 
    Showaxeslabels Bool `json:"showaxeslabels,omitempty"`
    
    // Showbackground 
    // boolean 
    // Sets whether or not this axis' wall has a background color. 
    Showbackground Bool `json:"showbackground,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent LayoutSceneXaxisShowexponent `json:"showexponent,omitempty"`
    
    // Showgrid 
    // boolean 
    // Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Showline 
    // boolean 
    // Determines whether or not a line bounding this axis is drawn. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showspikes 
    // boolean 
    // Sets whether or not spikes starting from data points to this axis' wall are shown on hover. 
    Showspikes Bool `json:"showspikes,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutSceneXaxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutSceneXaxisShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Spikecolor 
    // color 
    // Sets the color of the spikes. 
    Spikecolor Color `json:"spikecolor,omitempty"`
    
    // Spikesides 
    // boolean 
    // Sets whether or not spikes extending from the projection data points to this axis' wall boundaries are shown on hover. 
    Spikesides Bool `json:"spikesides,omitempty"`
    
    // Spikethickness 
    // number 
    // Sets the thickness (in px) of the spikes. 
    Spikethickness float64 `json:"spikethickness,omitempty"`
    
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
    //  Sets the tick font. 
    Tickfont *LayoutSceneXaxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutSceneXaxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutSceneXaxisTicks `json:"ticks,omitempty"`
    
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
    Title *LayoutSceneXaxisTitle `json:"title,omitempty"`
    
    // Type 
    // enumerated Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question. 
    Type LayoutSceneXaxisType `json:"type,omitempty"`
    
    // Visible 
    // boolean 
    // A single toggle to hide the axis while preserving interaction like dragging. Default is true when a cheater plot is present on the axis, otherwise false 
    Visible Bool `json:"visible,omitempty"`
    
    // Zeroline 
    // boolean 
    // Determines whether or not a line is drawn at along the 0 value of this axis. If *true*, the zero line is drawn on top of the grid lines. 
    Zeroline Bool `json:"zeroline,omitempty"`
    
    // Zerolinecolor 
    // color 
    // Sets the line color of the zero line. 
    Zerolinecolor Color `json:"zerolinecolor,omitempty"`
    
    // Zerolinewidth 
    // number 
    // Sets the width (in px) of the zero line. 
    Zerolinewidth float64 `json:"zerolinewidth,omitempty"`
    
}
// LayoutSceneYaxisTickfont Sets the tick font.
type LayoutSceneYaxisTickfont struct {
    
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
// LayoutSceneYaxisTitleFont Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute.
type LayoutSceneYaxisTitleFont struct {
    
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
// LayoutSceneYaxisTitle 
type LayoutSceneYaxisTitle struct {
    
    // Font 
    //  Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute. 
    Font *LayoutSceneYaxisTitleFont `json:"font,omitempty"`
    
    // Text 
    // string 
    // Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// LayoutSceneYaxis 
type LayoutSceneYaxis struct {
    
    // Autorange 
    // enumerated Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*. 
    Autorange LayoutSceneYaxisAutorange `json:"autorange,omitempty"`
    
    // Autotypenumbers 
    // enumerated Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers. 
    Autotypenumbers LayoutSceneYaxisAutotypenumbers `json:"autotypenumbers,omitempty"`
    
    // Backgroundcolor 
    // color 
    // Sets the background color of this axis' wall. 
    Backgroundcolor Color `json:"backgroundcolor,omitempty"`
    
    // Calendar 
    // enumerated Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar` 
    Calendar LayoutSceneYaxisCalendar `json:"calendar,omitempty"`
    
    // Categoryarray 
    // data_array 
    // Sets the order in which categories on this axis appear. Only has an effect if `categoryorder` is set to *array*. Used with `categoryorder`. 
    Categoryarray interface{} `json:"categoryarray,omitempty"`
    
    // Categoryarraysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  categoryarray . 
    Categoryarraysrc String `json:"categoryarraysrc,omitempty"`
    
    // Categoryorder 
    // enumerated Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values. 
    Categoryorder LayoutSceneYaxisCategoryorder `json:"categoryorder,omitempty"`
    
    // Color 
    // color 
    // Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this. 
    Color Color `json:"color,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat LayoutSceneYaxisExponentformat `json:"exponentformat,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the color of the grid lines. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the width (in px) of the grid lines. 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Hoverformat 
    // string 
    // Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Hoverformat String `json:"hoverformat,omitempty"`
    
    // Linecolor 
    // color 
    // Sets the axis line color. 
    Linecolor Color `json:"linecolor,omitempty"`
    
    // Linewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Linewidth float64 `json:"linewidth,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Mirror 
    // enumerated Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots. 
    Mirror LayoutSceneYaxisMirror `json:"mirror,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Range 
    // info_array 
    // Sets the range of this axis. If the axis `type` is *log*, then you must take the log of your desired range (e.g. to set the range from 1 to 100, set the range from 0 to 2). If the axis `type` is *date*, it should be date strings, like date data, though Date objects and unix milliseconds will be accepted and converted to strings. If the axis `type` is *category*, it should be numbers, using the scale where each category is assigned a serial number from zero in the order it appears. 
    Range interface{} `json:"range,omitempty"`
    
    // Rangemode 
    // enumerated If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes. 
    Rangemode LayoutSceneYaxisRangemode `json:"rangemode,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showaxeslabels 
    // boolean 
    // Sets whether or not this axis is labeled 
    Showaxeslabels Bool `json:"showaxeslabels,omitempty"`
    
    // Showbackground 
    // boolean 
    // Sets whether or not this axis' wall has a background color. 
    Showbackground Bool `json:"showbackground,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent LayoutSceneYaxisShowexponent `json:"showexponent,omitempty"`
    
    // Showgrid 
    // boolean 
    // Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Showline 
    // boolean 
    // Determines whether or not a line bounding this axis is drawn. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showspikes 
    // boolean 
    // Sets whether or not spikes starting from data points to this axis' wall are shown on hover. 
    Showspikes Bool `json:"showspikes,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutSceneYaxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutSceneYaxisShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Spikecolor 
    // color 
    // Sets the color of the spikes. 
    Spikecolor Color `json:"spikecolor,omitempty"`
    
    // Spikesides 
    // boolean 
    // Sets whether or not spikes extending from the projection data points to this axis' wall boundaries are shown on hover. 
    Spikesides Bool `json:"spikesides,omitempty"`
    
    // Spikethickness 
    // number 
    // Sets the thickness (in px) of the spikes. 
    Spikethickness float64 `json:"spikethickness,omitempty"`
    
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
    //  Sets the tick font. 
    Tickfont *LayoutSceneYaxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutSceneYaxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutSceneYaxisTicks `json:"ticks,omitempty"`
    
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
    Title *LayoutSceneYaxisTitle `json:"title,omitempty"`
    
    // Type 
    // enumerated Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question. 
    Type LayoutSceneYaxisType `json:"type,omitempty"`
    
    // Visible 
    // boolean 
    // A single toggle to hide the axis while preserving interaction like dragging. Default is true when a cheater plot is present on the axis, otherwise false 
    Visible Bool `json:"visible,omitempty"`
    
    // Zeroline 
    // boolean 
    // Determines whether or not a line is drawn at along the 0 value of this axis. If *true*, the zero line is drawn on top of the grid lines. 
    Zeroline Bool `json:"zeroline,omitempty"`
    
    // Zerolinecolor 
    // color 
    // Sets the line color of the zero line. 
    Zerolinecolor Color `json:"zerolinecolor,omitempty"`
    
    // Zerolinewidth 
    // number 
    // Sets the width (in px) of the zero line. 
    Zerolinewidth float64 `json:"zerolinewidth,omitempty"`
    
}
// LayoutSceneZaxisTickfont Sets the tick font.
type LayoutSceneZaxisTickfont struct {
    
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
// LayoutSceneZaxisTitleFont Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute.
type LayoutSceneZaxisTitleFont struct {
    
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
// LayoutSceneZaxisTitle 
type LayoutSceneZaxisTitle struct {
    
    // Font 
    //  Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute. 
    Font *LayoutSceneZaxisTitleFont `json:"font,omitempty"`
    
    // Text 
    // string 
    // Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// LayoutSceneZaxis 
type LayoutSceneZaxis struct {
    
    // Autorange 
    // enumerated Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*. 
    Autorange LayoutSceneZaxisAutorange `json:"autorange,omitempty"`
    
    // Autotypenumbers 
    // enumerated Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers. 
    Autotypenumbers LayoutSceneZaxisAutotypenumbers `json:"autotypenumbers,omitempty"`
    
    // Backgroundcolor 
    // color 
    // Sets the background color of this axis' wall. 
    Backgroundcolor Color `json:"backgroundcolor,omitempty"`
    
    // Calendar 
    // enumerated Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar` 
    Calendar LayoutSceneZaxisCalendar `json:"calendar,omitempty"`
    
    // Categoryarray 
    // data_array 
    // Sets the order in which categories on this axis appear. Only has an effect if `categoryorder` is set to *array*. Used with `categoryorder`. 
    Categoryarray interface{} `json:"categoryarray,omitempty"`
    
    // Categoryarraysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  categoryarray . 
    Categoryarraysrc String `json:"categoryarraysrc,omitempty"`
    
    // Categoryorder 
    // enumerated Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values. 
    Categoryorder LayoutSceneZaxisCategoryorder `json:"categoryorder,omitempty"`
    
    // Color 
    // color 
    // Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this. 
    Color Color `json:"color,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat LayoutSceneZaxisExponentformat `json:"exponentformat,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the color of the grid lines. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the width (in px) of the grid lines. 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Hoverformat 
    // string 
    // Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Hoverformat String `json:"hoverformat,omitempty"`
    
    // Linecolor 
    // color 
    // Sets the axis line color. 
    Linecolor Color `json:"linecolor,omitempty"`
    
    // Linewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Linewidth float64 `json:"linewidth,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Mirror 
    // enumerated Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots. 
    Mirror LayoutSceneZaxisMirror `json:"mirror,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Range 
    // info_array 
    // Sets the range of this axis. If the axis `type` is *log*, then you must take the log of your desired range (e.g. to set the range from 1 to 100, set the range from 0 to 2). If the axis `type` is *date*, it should be date strings, like date data, though Date objects and unix milliseconds will be accepted and converted to strings. If the axis `type` is *category*, it should be numbers, using the scale where each category is assigned a serial number from zero in the order it appears. 
    Range interface{} `json:"range,omitempty"`
    
    // Rangemode 
    // enumerated If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes. 
    Rangemode LayoutSceneZaxisRangemode `json:"rangemode,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showaxeslabels 
    // boolean 
    // Sets whether or not this axis is labeled 
    Showaxeslabels Bool `json:"showaxeslabels,omitempty"`
    
    // Showbackground 
    // boolean 
    // Sets whether or not this axis' wall has a background color. 
    Showbackground Bool `json:"showbackground,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent LayoutSceneZaxisShowexponent `json:"showexponent,omitempty"`
    
    // Showgrid 
    // boolean 
    // Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Showline 
    // boolean 
    // Determines whether or not a line bounding this axis is drawn. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showspikes 
    // boolean 
    // Sets whether or not spikes starting from data points to this axis' wall are shown on hover. 
    Showspikes Bool `json:"showspikes,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutSceneZaxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutSceneZaxisShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Spikecolor 
    // color 
    // Sets the color of the spikes. 
    Spikecolor Color `json:"spikecolor,omitempty"`
    
    // Spikesides 
    // boolean 
    // Sets whether or not spikes extending from the projection data points to this axis' wall boundaries are shown on hover. 
    Spikesides Bool `json:"spikesides,omitempty"`
    
    // Spikethickness 
    // number 
    // Sets the thickness (in px) of the spikes. 
    Spikethickness float64 `json:"spikethickness,omitempty"`
    
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
    //  Sets the tick font. 
    Tickfont *LayoutSceneZaxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutSceneZaxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutSceneZaxisTicks `json:"ticks,omitempty"`
    
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
    Title *LayoutSceneZaxisTitle `json:"title,omitempty"`
    
    // Type 
    // enumerated Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question. 
    Type LayoutSceneZaxisType `json:"type,omitempty"`
    
    // Visible 
    // boolean 
    // A single toggle to hide the axis while preserving interaction like dragging. Default is true when a cheater plot is present on the axis, otherwise false 
    Visible Bool `json:"visible,omitempty"`
    
    // Zeroline 
    // boolean 
    // Determines whether or not a line is drawn at along the 0 value of this axis. If *true*, the zero line is drawn on top of the grid lines. 
    Zeroline Bool `json:"zeroline,omitempty"`
    
    // Zerolinecolor 
    // color 
    // Sets the line color of the zero line. 
    Zerolinecolor Color `json:"zerolinecolor,omitempty"`
    
    // Zerolinewidth 
    // number 
    // Sets the width (in px) of the zero line. 
    Zerolinewidth float64 `json:"zerolinewidth,omitempty"`
    
}
// LayoutScene 
type LayoutScene struct {
    
    // Annotations 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Annotations interface{} `json:"annotations,omitempty"`
    
    // Aspectmode 
    // enumerated If *cube*, this scene's axes are drawn as a cube, regardless of the axes' ranges. If *data*, this scene's axes are drawn in proportion with the axes' ranges. If *manual*, this scene's axes are drawn in proportion with the input of *aspectratio* (the default behavior if *aspectratio* is provided). If *auto*, this scene's axes are drawn using the results of *data* except when one axis is more than four times the size of the two others, where in that case the results of *cube* are used. 
    Aspectmode LayoutSceneAspectmode `json:"aspectmode,omitempty"`
    
    // Aspectratio 
    //  Sets this scene's axis aspectratio. 
    Aspectratio *LayoutSceneAspectratio `json:"aspectratio,omitempty"`
    
    // Bgcolor 
    // color 
    //  
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Camera 
    //   
    Camera *LayoutSceneCamera `json:"camera,omitempty"`
    
    // Domain 
    //   
    Domain *LayoutSceneDomain `json:"domain,omitempty"`
    
    // Dragmode 
    // enumerated Determines the mode of drag interactions for this scene. 
    Dragmode LayoutSceneDragmode `json:"dragmode,omitempty"`
    
    // Hovermode 
    // enumerated Determines the mode of hover interactions for this scene. 
    Hovermode LayoutSceneHovermode `json:"hovermode,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in camera attributes. Defaults to `layout.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Xaxis 
    //   
    Xaxis *LayoutSceneXaxis `json:"xaxis,omitempty"`
    
    // Yaxis 
    //   
    Yaxis *LayoutSceneYaxis `json:"yaxis,omitempty"`
    
    // Zaxis 
    //   
    Zaxis *LayoutSceneZaxis `json:"zaxis,omitempty"`
    
}
// LayoutTernaryAaxisTickfont Sets the tick font.
type LayoutTernaryAaxisTickfont struct {
    
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
// LayoutTernaryAaxisTitleFont Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute.
type LayoutTernaryAaxisTitleFont struct {
    
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
// LayoutTernaryAaxisTitle 
type LayoutTernaryAaxisTitle struct {
    
    // Font 
    //  Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute. 
    Font *LayoutTernaryAaxisTitleFont `json:"font,omitempty"`
    
    // Text 
    // string 
    // Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// LayoutTernaryAaxis 
type LayoutTernaryAaxis struct {
    
    // Color 
    // color 
    // Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this. 
    Color Color `json:"color,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat LayoutTernaryAaxisExponentformat `json:"exponentformat,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the color of the grid lines. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the width (in px) of the grid lines. 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Hoverformat 
    // string 
    // Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Hoverformat String `json:"hoverformat,omitempty"`
    
    // Layer 
    // enumerated Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis. 
    Layer LayoutTernaryAaxisLayer `json:"layer,omitempty"`
    
    // Linecolor 
    // color 
    // Sets the axis line color. 
    Linecolor Color `json:"linecolor,omitempty"`
    
    // Linewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Linewidth float64 `json:"linewidth,omitempty"`
    
    // Min 
    // number 
    // The minimum value visible on this axis. The maximum is determined by the sum minus the minimum values of the other two axes. The full view corresponds to all the minima set to zero. 
    Min float64 `json:"min,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent LayoutTernaryAaxisShowexponent `json:"showexponent,omitempty"`
    
    // Showgrid 
    // boolean 
    // Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Showline 
    // boolean 
    // Determines whether or not a line bounding this axis is drawn. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutTernaryAaxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutTernaryAaxisShowticksuffix `json:"showticksuffix,omitempty"`
    
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
    //  Sets the tick font. 
    Tickfont *LayoutTernaryAaxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutTernaryAaxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutTernaryAaxisTicks `json:"ticks,omitempty"`
    
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
    Title *LayoutTernaryAaxisTitle `json:"title,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in axis `min`, and `title` if in `editable: true` configuration. Defaults to `ternary<N>.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
}
// LayoutTernaryBaxisTickfont Sets the tick font.
type LayoutTernaryBaxisTickfont struct {
    
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
// LayoutTernaryBaxisTitleFont Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute.
type LayoutTernaryBaxisTitleFont struct {
    
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
// LayoutTernaryBaxisTitle 
type LayoutTernaryBaxisTitle struct {
    
    // Font 
    //  Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute. 
    Font *LayoutTernaryBaxisTitleFont `json:"font,omitempty"`
    
    // Text 
    // string 
    // Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// LayoutTernaryBaxis 
type LayoutTernaryBaxis struct {
    
    // Color 
    // color 
    // Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this. 
    Color Color `json:"color,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat LayoutTernaryBaxisExponentformat `json:"exponentformat,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the color of the grid lines. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the width (in px) of the grid lines. 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Hoverformat 
    // string 
    // Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Hoverformat String `json:"hoverformat,omitempty"`
    
    // Layer 
    // enumerated Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis. 
    Layer LayoutTernaryBaxisLayer `json:"layer,omitempty"`
    
    // Linecolor 
    // color 
    // Sets the axis line color. 
    Linecolor Color `json:"linecolor,omitempty"`
    
    // Linewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Linewidth float64 `json:"linewidth,omitempty"`
    
    // Min 
    // number 
    // The minimum value visible on this axis. The maximum is determined by the sum minus the minimum values of the other two axes. The full view corresponds to all the minima set to zero. 
    Min float64 `json:"min,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent LayoutTernaryBaxisShowexponent `json:"showexponent,omitempty"`
    
    // Showgrid 
    // boolean 
    // Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Showline 
    // boolean 
    // Determines whether or not a line bounding this axis is drawn. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutTernaryBaxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutTernaryBaxisShowticksuffix `json:"showticksuffix,omitempty"`
    
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
    //  Sets the tick font. 
    Tickfont *LayoutTernaryBaxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutTernaryBaxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutTernaryBaxisTicks `json:"ticks,omitempty"`
    
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
    Title *LayoutTernaryBaxisTitle `json:"title,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in axis `min`, and `title` if in `editable: true` configuration. Defaults to `ternary<N>.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
}
// LayoutTernaryCaxisTickfont Sets the tick font.
type LayoutTernaryCaxisTickfont struct {
    
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
// LayoutTernaryCaxisTitleFont Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute.
type LayoutTernaryCaxisTitleFont struct {
    
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
// LayoutTernaryCaxisTitle 
type LayoutTernaryCaxisTitle struct {
    
    // Font 
    //  Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute. 
    Font *LayoutTernaryCaxisTitleFont `json:"font,omitempty"`
    
    // Text 
    // string 
    // Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// LayoutTernaryCaxis 
type LayoutTernaryCaxis struct {
    
    // Color 
    // color 
    // Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this. 
    Color Color `json:"color,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat LayoutTernaryCaxisExponentformat `json:"exponentformat,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the color of the grid lines. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the width (in px) of the grid lines. 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Hoverformat 
    // string 
    // Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Hoverformat String `json:"hoverformat,omitempty"`
    
    // Layer 
    // enumerated Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis. 
    Layer LayoutTernaryCaxisLayer `json:"layer,omitempty"`
    
    // Linecolor 
    // color 
    // Sets the axis line color. 
    Linecolor Color `json:"linecolor,omitempty"`
    
    // Linewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Linewidth float64 `json:"linewidth,omitempty"`
    
    // Min 
    // number 
    // The minimum value visible on this axis. The maximum is determined by the sum minus the minimum values of the other two axes. The full view corresponds to all the minima set to zero. 
    Min float64 `json:"min,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent LayoutTernaryCaxisShowexponent `json:"showexponent,omitempty"`
    
    // Showgrid 
    // boolean 
    // Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Showline 
    // boolean 
    // Determines whether or not a line bounding this axis is drawn. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutTernaryCaxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutTernaryCaxisShowticksuffix `json:"showticksuffix,omitempty"`
    
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
    //  Sets the tick font. 
    Tickfont *LayoutTernaryCaxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutTernaryCaxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutTernaryCaxisTicks `json:"ticks,omitempty"`
    
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
    Title *LayoutTernaryCaxisTitle `json:"title,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in axis `min`, and `title` if in `editable: true` configuration. Defaults to `ternary<N>.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
}
// LayoutTernaryDomain 
type LayoutTernaryDomain struct {
    
    // Column 
    // integer 
    // If there is a layout grid, use the domain for this column in the grid for this ternary subplot . 
    Column int64 `json:"column,omitempty"`
    
    // Row 
    // integer 
    // If there is a layout grid, use the domain for this row in the grid for this ternary subplot . 
    Row int64 `json:"row,omitempty"`
    
    // X 
    // info_array 
    // Sets the horizontal domain of this ternary subplot (in plot fraction). 
    X interface{} `json:"x,omitempty"`
    
    // Y 
    // info_array 
    // Sets the vertical domain of this ternary subplot (in plot fraction). 
    Y interface{} `json:"y,omitempty"`
    
}
// LayoutTernary 
type LayoutTernary struct {
    
    // Aaxis 
    //   
    Aaxis *LayoutTernaryAaxis `json:"aaxis,omitempty"`
    
    // Baxis 
    //   
    Baxis *LayoutTernaryBaxis `json:"baxis,omitempty"`
    
    // Bgcolor 
    // color 
    // Set the background color of the subplot 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Caxis 
    //   
    Caxis *LayoutTernaryCaxis `json:"caxis,omitempty"`
    
    // Domain 
    //   
    Domain *LayoutTernaryDomain `json:"domain,omitempty"`
    
    // Sum 
    // number 
    // The number each triplet should sum to, and the maximum range of each axis 
    Sum float64 `json:"sum,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in axis `min` and `title`, if not overridden in the individual axes. Defaults to `layout.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
}
// LayoutTitleFont Sets the title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute.
type LayoutTitleFont struct {
    
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
// LayoutTitlePad Sets the padding of the title. Each padding value only applies when the corresponding `xanchor`/`yanchor` value is set accordingly. E.g. for left padding to take effect, `xanchor` must be set to *left*. The same rule applies if `xanchor`/`yanchor` is determined automatically. Padding is muted if the respective anchor value is *middle*/*center*.
type LayoutTitlePad struct {
    
    // B 
    // number 
    // The amount of padding (in px) along the bottom of the component. 
    B float64 `json:"b,omitempty"`
    
    // L 
    // number 
    // The amount of padding (in px) on the left side of the component. 
    L float64 `json:"l,omitempty"`
    
    // R 
    // number 
    // The amount of padding (in px) on the right side of the component. 
    R float64 `json:"r,omitempty"`
    
    // T 
    // number 
    // The amount of padding (in px) along the top of the component. 
    T float64 `json:"t,omitempty"`
    
}
// LayoutTitle 
type LayoutTitle struct {
    
    // Font 
    //  Sets the title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute. 
    Font *LayoutTitleFont `json:"font,omitempty"`
    
    // Pad 
    //  Sets the padding of the title. Each padding value only applies when the corresponding `xanchor`/`yanchor` value is set accordingly. E.g. for left padding to take effect, `xanchor` must be set to *left*. The same rule applies if `xanchor`/`yanchor` is determined automatically. Padding is muted if the respective anchor value is *middle*/*center*. 
    Pad *LayoutTitlePad `json:"pad,omitempty"`
    
    // Text 
    // string 
    // Sets the plot's title. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
    // X 
    // number 
    // Sets the x position with respect to `xref` in normalized coordinates from *0* (left) to *1* (right). 
    X float64 `json:"x,omitempty"`
    
    // Xanchor 
    // enumerated Sets the title's horizontal alignment with respect to its x position. *left* means that the title starts at x, *right* means that the title ends at x and *center* means that the title's center is at x. *auto* divides `xref` by three and calculates the `xanchor` value automatically based on the value of `x`. 
    Xanchor LayoutTitleXanchor `json:"xanchor,omitempty"`
    
    // Xref 
    // enumerated Sets the container `x` refers to. *container* spans the entire `width` of the plot. *paper* refers to the width of the plotting area only. 
    Xref LayoutTitleXref `json:"xref,omitempty"`
    
    // Y 
    // number 
    // Sets the y position with respect to `yref` in normalized coordinates from *0* (bottom) to *1* (top). *auto* places the baseline of the title onto the vertical center of the top margin. 
    Y float64 `json:"y,omitempty"`
    
    // Yanchor 
    // enumerated Sets the title's vertical alignment with respect to its y position. *top* means that the title's cap line is at y, *bottom* means that the title's baseline is at y and *middle* means that the title's midline is at y. *auto* divides `yref` by three and calculates the `yanchor` value automatically based on the value of `y`. 
    Yanchor LayoutTitleYanchor `json:"yanchor,omitempty"`
    
    // Yref 
    // enumerated Sets the container `y` refers to. *container* spans the entire `height` of the plot. *paper* refers to the height of the plotting area only. 
    Yref LayoutTitleYref `json:"yref,omitempty"`
    
}
// LayoutTransition Sets transition options used during Plotly.react updates.
type LayoutTransition struct {
    
    // Duration 
    // number 
    // The duration of the transition, in milliseconds. If equal to zero, updates are synchronous. 
    Duration float64 `json:"duration,omitempty"`
    
    // Easing 
    // enumerated The easing function used for the transition 
    Easing LayoutTransitionEasing `json:"easing,omitempty"`
    
    // Ordering 
    // enumerated Determines whether the figure's layout or traces smoothly transitions during updates that make both traces and layout change. 
    Ordering LayoutTransitionOrdering `json:"ordering,omitempty"`
    
}
// LayoutUniformtext 
type LayoutUniformtext struct {
    
    // Minsize 
    // number 
    // Sets the minimum text size between traces of the same type. 
    Minsize float64 `json:"minsize,omitempty"`
    
    // Mode 
    // enumerated Determines how the font size for various text elements are uniformed between each trace type. If the computed text sizes were smaller than the minimum size defined by `uniformtext.minsize` using *hide* option hides the text; and using *show* option shows the text without further downscaling. Please note that if the size defined by `minsize` is greater than the font size defined by trace, then the `minsize` is used. 
    Mode LayoutUniformtextMode `json:"mode,omitempty"`
    
}
// LayoutXaxisRangeselectorFont Sets the font of the range selector button text.
type LayoutXaxisRangeselectorFont struct {
    
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
// LayoutXaxisRangeselector 
type LayoutXaxisRangeselector struct {
    
    // Activecolor 
    // color 
    // Sets the background color of the active range selector button. 
    Activecolor Color `json:"activecolor,omitempty"`
    
    // Bgcolor 
    // color 
    // Sets the background color of the range selector buttons. 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Bordercolor 
    // color 
    // Sets the color of the border enclosing the range selector. 
    Bordercolor Color `json:"bordercolor,omitempty"`
    
    // Borderwidth 
    // number 
    // Sets the width (in px) of the border enclosing the range selector. 
    Borderwidth float64 `json:"borderwidth,omitempty"`
    
    // Buttons 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Buttons interface{} `json:"buttons,omitempty"`
    
    // Font 
    //  Sets the font of the range selector button text. 
    Font *LayoutXaxisRangeselectorFont `json:"font,omitempty"`
    
    // Visible 
    // boolean 
    // Determines whether or not this range selector is visible. Note that range selectors are only available for x axes of `type` set to or auto-typed to *date*. 
    Visible Bool `json:"visible,omitempty"`
    
    // X 
    // number 
    // Sets the x position (in normalized coordinates) of the range selector. 
    X float64 `json:"x,omitempty"`
    
    // Xanchor 
    // enumerated Sets the range selector's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the range selector. 
    Xanchor LayoutXaxisRangeselectorXanchor `json:"xanchor,omitempty"`
    
    // Y 
    // number 
    // Sets the y position (in normalized coordinates) of the range selector. 
    Y float64 `json:"y,omitempty"`
    
    // Yanchor 
    // enumerated Sets the range selector's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the range selector. 
    Yanchor LayoutXaxisRangeselectorYanchor `json:"yanchor,omitempty"`
    
}
// LayoutXaxisRangesliderYaxis 
type LayoutXaxisRangesliderYaxis struct {
    
    // Range 
    // info_array 
    // Sets the range of this axis for the rangeslider. 
    Range interface{} `json:"range,omitempty"`
    
    // Rangemode 
    // enumerated Determines whether or not the range of this axis in the rangeslider use the same value than in the main plot when zooming in/out. If *auto*, the autorange will be used. If *fixed*, the `range` is used. If *match*, the current range of the corresponding y-axis on the main subplot is used. 
    Rangemode LayoutXaxisRangesliderYaxisRangemode `json:"rangemode,omitempty"`
    
}
// LayoutXaxisRangeslider 
type LayoutXaxisRangeslider struct {
    
    // Autorange 
    // boolean 
    // Determines whether or not the range slider range is computed in relation to the input data. If `range` is provided, then `autorange` is set to *false*. 
    Autorange Bool `json:"autorange,omitempty"`
    
    // Bgcolor 
    // color 
    // Sets the background color of the range slider. 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Bordercolor 
    // color 
    // Sets the border color of the range slider. 
    Bordercolor Color `json:"bordercolor,omitempty"`
    
    // Borderwidth 
    // integer 
    // Sets the border width of the range slider. 
    Borderwidth int64 `json:"borderwidth,omitempty"`
    
    // Range 
    // info_array 
    // Sets the range of the range slider. If not set, defaults to the full xaxis range. If the axis `type` is *log*, then you must take the log of your desired range. If the axis `type` is *date*, it should be date strings, like date data, though Date objects and unix milliseconds will be accepted and converted to strings. If the axis `type` is *category*, it should be numbers, using the scale where each category is assigned a serial number from zero in the order it appears. 
    Range interface{} `json:"range,omitempty"`
    
    // Thickness 
    // number 
    // The height of the range slider as a fraction of the total plot area height. 
    Thickness float64 `json:"thickness,omitempty"`
    
    // Visible 
    // boolean 
    // Determines whether or not the range slider will be visible. If visible, perpendicular axes will be set to `fixedrange` 
    Visible Bool `json:"visible,omitempty"`
    
    // Yaxis 
    //   
    Yaxis *LayoutXaxisRangesliderYaxis `json:"yaxis,omitempty"`
    
}
// LayoutXaxisTickfont Sets the tick font.
type LayoutXaxisTickfont struct {
    
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
// LayoutXaxisTitleFont Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute.
type LayoutXaxisTitleFont struct {
    
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
// LayoutXaxisTitle 
type LayoutXaxisTitle struct {
    
    // Font 
    //  Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute. 
    Font *LayoutXaxisTitleFont `json:"font,omitempty"`
    
    // Standoff 
    // number 
    // Sets the standoff distance (in px) between the axis labels and the title text The default value is a function of the axis tick labels, the title `font.size` and the axis `linewidth`. Note that the axis title position is always constrained within the margins, so the actual standoff distance is always less than the set or default value. By setting `standoff` and turning on `automargin`, plotly.js will push the margins to fit the axis title at given standoff distance. 
    Standoff float64 `json:"standoff,omitempty"`
    
    // Text 
    // string 
    // Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// LayoutXaxis 
type LayoutXaxis struct {
    
    // Anchor 
    // enumerated If set to an opposite-letter axis id (e.g. `x2`, `y`), this axis is bound to the corresponding opposite-letter axis. If set to *free*, this axis' position is determined by `position`. 
    Anchor LayoutXaxisAnchor `json:"anchor,omitempty"`
    
    // Automargin 
    // boolean 
    // Determines whether long tick labels automatically grow the figure margins. 
    Automargin Bool `json:"automargin,omitempty"`
    
    // Autorange 
    // enumerated Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*. 
    Autorange LayoutXaxisAutorange `json:"autorange,omitempty"`
    
    // Autotypenumbers 
    // enumerated Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers. 
    Autotypenumbers LayoutXaxisAutotypenumbers `json:"autotypenumbers,omitempty"`
    
    // Calendar 
    // enumerated Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar` 
    Calendar LayoutXaxisCalendar `json:"calendar,omitempty"`
    
    // Categoryarray 
    // data_array 
    // Sets the order in which categories on this axis appear. Only has an effect if `categoryorder` is set to *array*. Used with `categoryorder`. 
    Categoryarray interface{} `json:"categoryarray,omitempty"`
    
    // Categoryarraysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  categoryarray . 
    Categoryarraysrc String `json:"categoryarraysrc,omitempty"`
    
    // Categoryorder 
    // enumerated Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values. 
    Categoryorder LayoutXaxisCategoryorder `json:"categoryorder,omitempty"`
    
    // Color 
    // color 
    // Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this. 
    Color Color `json:"color,omitempty"`
    
    // Constrain 
    // enumerated If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines how that happens: by increasing the *range*, or by decreasing the *domain*. Default is *domain* for axes containing image traces, *range* otherwise. 
    Constrain LayoutXaxisConstrain `json:"constrain,omitempty"`
    
    // Constraintoward 
    // enumerated If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines which direction we push the originally specified plot area. Options are *left*, *center* (default), and *right* for x axes, and *top*, *middle* (default), and *bottom* for y axes. 
    Constraintoward LayoutXaxisConstraintoward `json:"constraintoward,omitempty"`
    
    // Dividercolor 
    // color 
    // Sets the color of the dividers Only has an effect on *multicategory* axes. 
    Dividercolor Color `json:"dividercolor,omitempty"`
    
    // Dividerwidth 
    // number 
    // Sets the width (in px) of the dividers Only has an effect on *multicategory* axes. 
    Dividerwidth float64 `json:"dividerwidth,omitempty"`
    
    // Domain 
    // info_array 
    // Sets the domain of this axis (in plot fraction). 
    Domain interface{} `json:"domain,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat LayoutXaxisExponentformat `json:"exponentformat,omitempty"`
    
    // Fixedrange 
    // boolean 
    // Determines whether or not this axis is zoom-able. If true, then zoom is disabled. 
    Fixedrange Bool `json:"fixedrange,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the color of the grid lines. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the width (in px) of the grid lines. 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Hoverformat 
    // string 
    // Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Hoverformat String `json:"hoverformat,omitempty"`
    
    // Layer 
    // enumerated Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis. 
    Layer LayoutXaxisLayer `json:"layer,omitempty"`
    
    // Linecolor 
    // color 
    // Sets the axis line color. 
    Linecolor Color `json:"linecolor,omitempty"`
    
    // Linewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Linewidth float64 `json:"linewidth,omitempty"`
    
    // Matches 
    // enumerated If set to another axis id (e.g. `x2`, `y`), the range of this axis will match the range of the corresponding axis in data-coordinates space. Moreover, matching axes share auto-range values, category lists and histogram auto-bins. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden. Moreover, note that matching axes must have the same `type`. 
    Matches LayoutXaxisMatches `json:"matches,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Mirror 
    // enumerated Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots. 
    Mirror LayoutXaxisMirror `json:"mirror,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Overlaying 
    // enumerated If set a same-letter axis id, this axis is overlaid on top of the corresponding same-letter axis, with traces and axes visible for both axes. If *false*, this axis does not overlay any same-letter axes. In this case, for axes with overlapping domains only the highest-numbered axis will be visible. 
    Overlaying LayoutXaxisOverlaying `json:"overlaying,omitempty"`
    
    // Position 
    // number 
    // Sets the position of this axis in the plotting space (in normalized coordinates). Only has an effect if `anchor` is set to *free*. 
    Position float64 `json:"position,omitempty"`
    
    // Range 
    // info_array 
    // Sets the range of this axis. If the axis `type` is *log*, then you must take the log of your desired range (e.g. to set the range from 1 to 100, set the range from 0 to 2). If the axis `type` is *date*, it should be date strings, like date data, though Date objects and unix milliseconds will be accepted and converted to strings. If the axis `type` is *category*, it should be numbers, using the scale where each category is assigned a serial number from zero in the order it appears. 
    Range interface{} `json:"range,omitempty"`
    
    // Rangebreaks 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Rangebreaks interface{} `json:"rangebreaks,omitempty"`
    
    // Rangemode 
    // enumerated If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes. 
    Rangemode LayoutXaxisRangemode `json:"rangemode,omitempty"`
    
    // Rangeselector 
    //   
    Rangeselector *LayoutXaxisRangeselector `json:"rangeselector,omitempty"`
    
    // Rangeslider 
    //   
    Rangeslider *LayoutXaxisRangeslider `json:"rangeslider,omitempty"`
    
    // Scaleanchor 
    // enumerated If set to another axis id (e.g. `x2`, `y`), the range of this axis changes together with the range of the corresponding axis such that the scale of pixels per unit is in a constant ratio. Both axes are still zoomable, but when you zoom one, the other will zoom the same amount, keeping a fixed midpoint. `constrain` and `constraintoward` determine how we enforce the constraint. You can chain these, ie `yaxis: {scaleanchor: *x*}, xaxis2: {scaleanchor: *y*}` but you can only link axes of the same `type`. The linked axis can have the opposite letter (to constrain the aspect ratio) or the same letter (to match scales across subplots). Loops (`yaxis: {scaleanchor: *x*}, xaxis: {scaleanchor: *y*}` or longer) are redundant and the last constraint encountered will be ignored to avoid possible inconsistent constraints via `scaleratio`. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden. 
    Scaleanchor LayoutXaxisScaleanchor `json:"scaleanchor,omitempty"`
    
    // Scaleratio 
    // number 
    // If this axis is linked to another by `scaleanchor`, this determines the pixel to unit scale ratio. For example, if this value is 10, then every unit on this axis spans 10 times the number of pixels as a unit on the linked axis. Use this for example to create an elevation profile where the vertical scale is exaggerated a fixed amount with respect to the horizontal. 
    Scaleratio float64 `json:"scaleratio,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showdividers 
    // boolean 
    // Determines whether or not a dividers are drawn between the category levels of this axis. Only has an effect on *multicategory* axes. 
    Showdividers Bool `json:"showdividers,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent LayoutXaxisShowexponent `json:"showexponent,omitempty"`
    
    // Showgrid 
    // boolean 
    // Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Showline 
    // boolean 
    // Determines whether or not a line bounding this axis is drawn. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showspikes 
    // boolean 
    // Determines whether or not spikes (aka droplines) are drawn for this axis. Note: This only takes affect when hovermode = closest 
    Showspikes Bool `json:"showspikes,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutXaxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutXaxisShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Side 
    // enumerated Determines whether a x (y) axis is positioned at the *bottom* (*left*) or *top* (*right*) of the plotting area. 
    Side LayoutXaxisSide `json:"side,omitempty"`
    
    // Spikecolor 
    // color 
    // Sets the spike color. If undefined, will use the series color 
    Spikecolor Color `json:"spikecolor,omitempty"`
    
    // Spikedash 
    // string 
    // Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*). 
    Spikedash String `json:"spikedash,omitempty"`
    
    // Spikemode 
    // flaglist Determines the drawing mode for the spike line If *toaxis*, the line is drawn from the data point to the axis the  series is plotted on. If *across*, the line is drawn across the entire plot area, and supercedes *toaxis*. If *marker*, then a marker dot is drawn on the axis the series is plotted on 
    Spikemode LayoutXaxisSpikemode `json:"spikemode,omitempty"`
    
    // Spikesnap 
    // enumerated Determines whether spikelines are stuck to the cursor or to the closest datapoints. 
    Spikesnap LayoutXaxisSpikesnap `json:"spikesnap,omitempty"`
    
    // Spikethickness 
    // number 
    // Sets the width (in px) of the zero line. 
    Spikethickness float64 `json:"spikethickness,omitempty"`
    
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
    //  Sets the tick font. 
    Tickfont *LayoutXaxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklabelmode 
    // enumerated Determines where tick labels are drawn with respect to their corresponding ticks and grid lines. Only has an effect for axes of `type` *date* When set to *period*, tick labels are drawn in the middle of the period between ticks. 
    Ticklabelmode LayoutXaxisTicklabelmode `json:"ticklabelmode,omitempty"`
    
    // Ticklabelposition 
    // enumerated Determines where tick labels are drawn with respect to the axis Please note that top or bottom has no effect on x axes or when `ticklabelmode` is set to *period*. Similarly left or right has no effect on y axes or when `ticklabelmode` is set to *period*. Has no effect on *multicategory* axes or when `tickson` is set to *boundaries*. When used on axes linked by `matches` or `scaleanchor`, no extra padding for inside labels would be added by autorange, so that the scales could match. 
    Ticklabelposition LayoutXaxisTicklabelposition `json:"ticklabelposition,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutXaxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutXaxisTicks `json:"ticks,omitempty"`
    
    // Tickson 
    // enumerated Determines where ticks and grid lines are drawn with respect to their corresponding tick labels. Only has an effect for axes of `type` *category* or *multicategory*. When set to *boundaries*, ticks and grid lines are drawn half a category to the left/bottom of labels. 
    Tickson LayoutXaxisTickson `json:"tickson,omitempty"`
    
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
    Title *LayoutXaxisTitle `json:"title,omitempty"`
    
    // Type 
    // enumerated Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question. 
    Type LayoutXaxisType `json:"type,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in axis `range`, `autorange`, and `title` if in `editable: true` configuration. Defaults to `layout.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Visible 
    // boolean 
    // A single toggle to hide the axis while preserving interaction like dragging. Default is true when a cheater plot is present on the axis, otherwise false 
    Visible Bool `json:"visible,omitempty"`
    
    // Zeroline 
    // boolean 
    // Determines whether or not a line is drawn at along the 0 value of this axis. If *true*, the zero line is drawn on top of the grid lines. 
    Zeroline Bool `json:"zeroline,omitempty"`
    
    // Zerolinecolor 
    // color 
    // Sets the line color of the zero line. 
    Zerolinecolor Color `json:"zerolinecolor,omitempty"`
    
    // Zerolinewidth 
    // number 
    // Sets the width (in px) of the zero line. 
    Zerolinewidth float64 `json:"zerolinewidth,omitempty"`
    
}
// LayoutYaxisTickfont Sets the tick font.
type LayoutYaxisTickfont struct {
    
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
// LayoutYaxisTitleFont Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute.
type LayoutYaxisTitleFont struct {
    
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
// LayoutYaxisTitle 
type LayoutYaxisTitle struct {
    
    // Font 
    //  Sets this axis' title font. Note that the title's font used to be customized by the now deprecated `titlefont` attribute. 
    Font *LayoutYaxisTitleFont `json:"font,omitempty"`
    
    // Standoff 
    // number 
    // Sets the standoff distance (in px) between the axis labels and the title text The default value is a function of the axis tick labels, the title `font.size` and the axis `linewidth`. Note that the axis title position is always constrained within the margins, so the actual standoff distance is always less than the set or default value. By setting `standoff` and turning on `automargin`, plotly.js will push the margins to fit the axis title at given standoff distance. 
    Standoff float64 `json:"standoff,omitempty"`
    
    // Text 
    // string 
    // Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// LayoutYaxis 
type LayoutYaxis struct {
    
    // Anchor 
    // enumerated If set to an opposite-letter axis id (e.g. `x2`, `y`), this axis is bound to the corresponding opposite-letter axis. If set to *free*, this axis' position is determined by `position`. 
    Anchor LayoutYaxisAnchor `json:"anchor,omitempty"`
    
    // Automargin 
    // boolean 
    // Determines whether long tick labels automatically grow the figure margins. 
    Automargin Bool `json:"automargin,omitempty"`
    
    // Autorange 
    // enumerated Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*. 
    Autorange LayoutYaxisAutorange `json:"autorange,omitempty"`
    
    // Autotypenumbers 
    // enumerated Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers. 
    Autotypenumbers LayoutYaxisAutotypenumbers `json:"autotypenumbers,omitempty"`
    
    // Calendar 
    // enumerated Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar` 
    Calendar LayoutYaxisCalendar `json:"calendar,omitempty"`
    
    // Categoryarray 
    // data_array 
    // Sets the order in which categories on this axis appear. Only has an effect if `categoryorder` is set to *array*. Used with `categoryorder`. 
    Categoryarray interface{} `json:"categoryarray,omitempty"`
    
    // Categoryarraysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  categoryarray . 
    Categoryarraysrc String `json:"categoryarraysrc,omitempty"`
    
    // Categoryorder 
    // enumerated Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values. 
    Categoryorder LayoutYaxisCategoryorder `json:"categoryorder,omitempty"`
    
    // Color 
    // color 
    // Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this. 
    Color Color `json:"color,omitempty"`
    
    // Constrain 
    // enumerated If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines how that happens: by increasing the *range*, or by decreasing the *domain*. Default is *domain* for axes containing image traces, *range* otherwise. 
    Constrain LayoutYaxisConstrain `json:"constrain,omitempty"`
    
    // Constraintoward 
    // enumerated If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines which direction we push the originally specified plot area. Options are *left*, *center* (default), and *right* for x axes, and *top*, *middle* (default), and *bottom* for y axes. 
    Constraintoward LayoutYaxisConstraintoward `json:"constraintoward,omitempty"`
    
    // Dividercolor 
    // color 
    // Sets the color of the dividers Only has an effect on *multicategory* axes. 
    Dividercolor Color `json:"dividercolor,omitempty"`
    
    // Dividerwidth 
    // number 
    // Sets the width (in px) of the dividers Only has an effect on *multicategory* axes. 
    Dividerwidth float64 `json:"dividerwidth,omitempty"`
    
    // Domain 
    // info_array 
    // Sets the domain of this axis (in plot fraction). 
    Domain interface{} `json:"domain,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat LayoutYaxisExponentformat `json:"exponentformat,omitempty"`
    
    // Fixedrange 
    // boolean 
    // Determines whether or not this axis is zoom-able. If true, then zoom is disabled. 
    Fixedrange Bool `json:"fixedrange,omitempty"`
    
    // Gridcolor 
    // color 
    // Sets the color of the grid lines. 
    Gridcolor Color `json:"gridcolor,omitempty"`
    
    // Gridwidth 
    // number 
    // Sets the width (in px) of the grid lines. 
    Gridwidth float64 `json:"gridwidth,omitempty"`
    
    // Hoverformat 
    // string 
    // Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Hoverformat String `json:"hoverformat,omitempty"`
    
    // Layer 
    // enumerated Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis. 
    Layer LayoutYaxisLayer `json:"layer,omitempty"`
    
    // Linecolor 
    // color 
    // Sets the axis line color. 
    Linecolor Color `json:"linecolor,omitempty"`
    
    // Linewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Linewidth float64 `json:"linewidth,omitempty"`
    
    // Matches 
    // enumerated If set to another axis id (e.g. `x2`, `y`), the range of this axis will match the range of the corresponding axis in data-coordinates space. Moreover, matching axes share auto-range values, category lists and histogram auto-bins. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden. Moreover, note that matching axes must have the same `type`. 
    Matches LayoutYaxisMatches `json:"matches,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Mirror 
    // enumerated Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots. 
    Mirror LayoutYaxisMirror `json:"mirror,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Overlaying 
    // enumerated If set a same-letter axis id, this axis is overlaid on top of the corresponding same-letter axis, with traces and axes visible for both axes. If *false*, this axis does not overlay any same-letter axes. In this case, for axes with overlapping domains only the highest-numbered axis will be visible. 
    Overlaying LayoutYaxisOverlaying `json:"overlaying,omitempty"`
    
    // Position 
    // number 
    // Sets the position of this axis in the plotting space (in normalized coordinates). Only has an effect if `anchor` is set to *free*. 
    Position float64 `json:"position,omitempty"`
    
    // Range 
    // info_array 
    // Sets the range of this axis. If the axis `type` is *log*, then you must take the log of your desired range (e.g. to set the range from 1 to 100, set the range from 0 to 2). If the axis `type` is *date*, it should be date strings, like date data, though Date objects and unix milliseconds will be accepted and converted to strings. If the axis `type` is *category*, it should be numbers, using the scale where each category is assigned a serial number from zero in the order it appears. 
    Range interface{} `json:"range,omitempty"`
    
    // Rangebreaks 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Rangebreaks interface{} `json:"rangebreaks,omitempty"`
    
    // Rangemode 
    // enumerated If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes. 
    Rangemode LayoutYaxisRangemode `json:"rangemode,omitempty"`
    
    // Scaleanchor 
    // enumerated If set to another axis id (e.g. `x2`, `y`), the range of this axis changes together with the range of the corresponding axis such that the scale of pixels per unit is in a constant ratio. Both axes are still zoomable, but when you zoom one, the other will zoom the same amount, keeping a fixed midpoint. `constrain` and `constraintoward` determine how we enforce the constraint. You can chain these, ie `yaxis: {scaleanchor: *x*}, xaxis2: {scaleanchor: *y*}` but you can only link axes of the same `type`. The linked axis can have the opposite letter (to constrain the aspect ratio) or the same letter (to match scales across subplots). Loops (`yaxis: {scaleanchor: *x*}, xaxis: {scaleanchor: *y*}` or longer) are redundant and the last constraint encountered will be ignored to avoid possible inconsistent constraints via `scaleratio`. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden. 
    Scaleanchor LayoutYaxisScaleanchor `json:"scaleanchor,omitempty"`
    
    // Scaleratio 
    // number 
    // If this axis is linked to another by `scaleanchor`, this determines the pixel to unit scale ratio. For example, if this value is 10, then every unit on this axis spans 10 times the number of pixels as a unit on the linked axis. Use this for example to create an elevation profile where the vertical scale is exaggerated a fixed amount with respect to the horizontal. 
    Scaleratio float64 `json:"scaleratio,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showdividers 
    // boolean 
    // Determines whether or not a dividers are drawn between the category levels of this axis. Only has an effect on *multicategory* axes. 
    Showdividers Bool `json:"showdividers,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent LayoutYaxisShowexponent `json:"showexponent,omitempty"`
    
    // Showgrid 
    // boolean 
    // Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark. 
    Showgrid Bool `json:"showgrid,omitempty"`
    
    // Showline 
    // boolean 
    // Determines whether or not a line bounding this axis is drawn. 
    Showline Bool `json:"showline,omitempty"`
    
    // Showspikes 
    // boolean 
    // Determines whether or not spikes (aka droplines) are drawn for this axis. Note: This only takes affect when hovermode = closest 
    Showspikes Bool `json:"showspikes,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix LayoutYaxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix LayoutYaxisShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Side 
    // enumerated Determines whether a x (y) axis is positioned at the *bottom* (*left*) or *top* (*right*) of the plotting area. 
    Side LayoutYaxisSide `json:"side,omitempty"`
    
    // Spikecolor 
    // color 
    // Sets the spike color. If undefined, will use the series color 
    Spikecolor Color `json:"spikecolor,omitempty"`
    
    // Spikedash 
    // string 
    // Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*). 
    Spikedash String `json:"spikedash,omitempty"`
    
    // Spikemode 
    // flaglist Determines the drawing mode for the spike line If *toaxis*, the line is drawn from the data point to the axis the  series is plotted on. If *across*, the line is drawn across the entire plot area, and supercedes *toaxis*. If *marker*, then a marker dot is drawn on the axis the series is plotted on 
    Spikemode LayoutYaxisSpikemode `json:"spikemode,omitempty"`
    
    // Spikesnap 
    // enumerated Determines whether spikelines are stuck to the cursor or to the closest datapoints. 
    Spikesnap LayoutYaxisSpikesnap `json:"spikesnap,omitempty"`
    
    // Spikethickness 
    // number 
    // Sets the width (in px) of the zero line. 
    Spikethickness float64 `json:"spikethickness,omitempty"`
    
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
    //  Sets the tick font. 
    Tickfont *LayoutYaxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklabelmode 
    // enumerated Determines where tick labels are drawn with respect to their corresponding ticks and grid lines. Only has an effect for axes of `type` *date* When set to *period*, tick labels are drawn in the middle of the period between ticks. 
    Ticklabelmode LayoutYaxisTicklabelmode `json:"ticklabelmode,omitempty"`
    
    // Ticklabelposition 
    // enumerated Determines where tick labels are drawn with respect to the axis Please note that top or bottom has no effect on x axes or when `ticklabelmode` is set to *period*. Similarly left or right has no effect on y axes or when `ticklabelmode` is set to *period*. Has no effect on *multicategory* axes or when `tickson` is set to *boundaries*. When used on axes linked by `matches` or `scaleanchor`, no extra padding for inside labels would be added by autorange, so that the scales could match. 
    Ticklabelposition LayoutYaxisTicklabelposition `json:"ticklabelposition,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode LayoutYaxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks LayoutYaxisTicks `json:"ticks,omitempty"`
    
    // Tickson 
    // enumerated Determines where ticks and grid lines are drawn with respect to their corresponding tick labels. Only has an effect for axes of `type` *category* or *multicategory*. When set to *boundaries*, ticks and grid lines are drawn half a category to the left/bottom of labels. 
    Tickson LayoutYaxisTickson `json:"tickson,omitempty"`
    
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
    Title *LayoutYaxisTitle `json:"title,omitempty"`
    
    // Type 
    // enumerated Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question. 
    Type LayoutYaxisType `json:"type,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of user-driven changes in axis `range`, `autorange`, and `title` if in `editable: true` configuration. Defaults to `layout.uirevision`. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Visible 
    // boolean 
    // A single toggle to hide the axis while preserving interaction like dragging. Default is true when a cheater plot is present on the axis, otherwise false 
    Visible Bool `json:"visible,omitempty"`
    
    // Zeroline 
    // boolean 
    // Determines whether or not a line is drawn at along the 0 value of this axis. If *true*, the zero line is drawn on top of the grid lines. 
    Zeroline Bool `json:"zeroline,omitempty"`
    
    // Zerolinecolor 
    // color 
    // Sets the line color of the zero line. 
    Zerolinecolor Color `json:"zerolinecolor,omitempty"`
    
    // Zerolinewidth 
    // number 
    // Sets the width (in px) of the zero line. 
    Zerolinewidth float64 `json:"zerolinewidth,omitempty"`
    
}
// LayoutAngularaxisTickorientation Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the orientation (from the paper perspective) of the angular axis tick labels.
type LayoutAngularaxisTickorientation string 

const (
    LayoutAngularaxisTickorientationHorizontal LayoutAngularaxisTickorientation = "horizontal"
    LayoutAngularaxisTickorientationVertical LayoutAngularaxisTickorientation = "vertical"
    
)
// LayoutAutotypenumbers Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. This is the default value; however it could be overridden for individual axes.
type LayoutAutotypenumbers string 

const (
    LayoutAutotypenumbersConvertTypes LayoutAutotypenumbers = "convert types"
    LayoutAutotypenumbersStrict LayoutAutotypenumbers = "strict"
    
)
// LayoutCalendar Sets the default calendar system to use for interpreting and displaying dates throughout the plot.
type LayoutCalendar string 

const (
    LayoutCalendarGregorian LayoutCalendar = "gregorian"
    LayoutCalendarChinese LayoutCalendar = "chinese"
    LayoutCalendarCoptic LayoutCalendar = "coptic"
    LayoutCalendarDiscworld LayoutCalendar = "discworld"
    LayoutCalendarEthiopian LayoutCalendar = "ethiopian"
    LayoutCalendarHebrew LayoutCalendar = "hebrew"
    LayoutCalendarIslamic LayoutCalendar = "islamic"
    LayoutCalendarJulian LayoutCalendar = "julian"
    LayoutCalendarMayan LayoutCalendar = "mayan"
    LayoutCalendarNanakshahi LayoutCalendar = "nanakshahi"
    LayoutCalendarNepali LayoutCalendar = "nepali"
    LayoutCalendarPersian LayoutCalendar = "persian"
    LayoutCalendarJalali LayoutCalendar = "jalali"
    LayoutCalendarTaiwan LayoutCalendar = "taiwan"
    LayoutCalendarThai LayoutCalendar = "thai"
    LayoutCalendarUmmalqura LayoutCalendar = "ummalqura"
    
)
// LayoutColoraxisColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutColoraxisColorbarExponentformat string 

const (
    LayoutColoraxisColorbarExponentformatNone LayoutColoraxisColorbarExponentformat = "none"
    LayoutColoraxisColorbarExponentformatE1 LayoutColoraxisColorbarExponentformat = "e"
    LayoutColoraxisColorbarExponentformatE2 LayoutColoraxisColorbarExponentformat = "E"
    LayoutColoraxisColorbarExponentformatPower LayoutColoraxisColorbarExponentformat = "power"
    LayoutColoraxisColorbarExponentformatSi LayoutColoraxisColorbarExponentformat = "SI"
    LayoutColoraxisColorbarExponentformatB LayoutColoraxisColorbarExponentformat = "B"
    
)
// LayoutColoraxisColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type LayoutColoraxisColorbarLenmode string 

const (
    LayoutColoraxisColorbarLenmodeFraction LayoutColoraxisColorbarLenmode = "fraction"
    LayoutColoraxisColorbarLenmodePixels LayoutColoraxisColorbarLenmode = "pixels"
    
)
// LayoutColoraxisColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutColoraxisColorbarShowexponent string 

const (
    LayoutColoraxisColorbarShowexponentAll LayoutColoraxisColorbarShowexponent = "all"
    LayoutColoraxisColorbarShowexponentFirst LayoutColoraxisColorbarShowexponent = "first"
    LayoutColoraxisColorbarShowexponentLast LayoutColoraxisColorbarShowexponent = "last"
    LayoutColoraxisColorbarShowexponentNone LayoutColoraxisColorbarShowexponent = "none"
    
)
// LayoutColoraxisColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutColoraxisColorbarShowtickprefix string 

const (
    LayoutColoraxisColorbarShowtickprefixAll LayoutColoraxisColorbarShowtickprefix = "all"
    LayoutColoraxisColorbarShowtickprefixFirst LayoutColoraxisColorbarShowtickprefix = "first"
    LayoutColoraxisColorbarShowtickprefixLast LayoutColoraxisColorbarShowtickprefix = "last"
    LayoutColoraxisColorbarShowtickprefixNone LayoutColoraxisColorbarShowtickprefix = "none"
    
)
// LayoutColoraxisColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutColoraxisColorbarShowticksuffix string 

const (
    LayoutColoraxisColorbarShowticksuffixAll LayoutColoraxisColorbarShowticksuffix = "all"
    LayoutColoraxisColorbarShowticksuffixFirst LayoutColoraxisColorbarShowticksuffix = "first"
    LayoutColoraxisColorbarShowticksuffixLast LayoutColoraxisColorbarShowticksuffix = "last"
    LayoutColoraxisColorbarShowticksuffixNone LayoutColoraxisColorbarShowticksuffix = "none"
    
)
// LayoutColoraxisColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type LayoutColoraxisColorbarThicknessmode string 

const (
    LayoutColoraxisColorbarThicknessmodeFraction LayoutColoraxisColorbarThicknessmode = "fraction"
    LayoutColoraxisColorbarThicknessmodePixels LayoutColoraxisColorbarThicknessmode = "pixels"
    
)
// LayoutColoraxisColorbarTicklabelposition Determines where tick labels are drawn.
type LayoutColoraxisColorbarTicklabelposition string 

const (
    LayoutColoraxisColorbarTicklabelpositionOutside LayoutColoraxisColorbarTicklabelposition = "outside"
    LayoutColoraxisColorbarTicklabelpositionInside LayoutColoraxisColorbarTicklabelposition = "inside"
    LayoutColoraxisColorbarTicklabelpositionOutsideTop LayoutColoraxisColorbarTicklabelposition = "outside top"
    LayoutColoraxisColorbarTicklabelpositionInsideTop LayoutColoraxisColorbarTicklabelposition = "inside top"
    LayoutColoraxisColorbarTicklabelpositionOutsideBottom LayoutColoraxisColorbarTicklabelposition = "outside bottom"
    LayoutColoraxisColorbarTicklabelpositionInsideBottom LayoutColoraxisColorbarTicklabelposition = "inside bottom"
    
)
// LayoutColoraxisColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutColoraxisColorbarTickmode string 

const (
    LayoutColoraxisColorbarTickmodeAuto LayoutColoraxisColorbarTickmode = "auto"
    LayoutColoraxisColorbarTickmodeLinear LayoutColoraxisColorbarTickmode = "linear"
    LayoutColoraxisColorbarTickmodeArray LayoutColoraxisColorbarTickmode = "array"
    
)
// LayoutColoraxisColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutColoraxisColorbarTicks string 

const (
    LayoutColoraxisColorbarTicksOutside LayoutColoraxisColorbarTicks = "outside"
    LayoutColoraxisColorbarTicksInside LayoutColoraxisColorbarTicks = "inside"
    LayoutColoraxisColorbarTicksEmpty LayoutColoraxisColorbarTicks = ""
    
)
// LayoutColoraxisColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type LayoutColoraxisColorbarTitleSide string 

const (
    LayoutColoraxisColorbarTitleSideRight LayoutColoraxisColorbarTitleSide = "right"
    LayoutColoraxisColorbarTitleSideTop LayoutColoraxisColorbarTitleSide = "top"
    LayoutColoraxisColorbarTitleSideBottom LayoutColoraxisColorbarTitleSide = "bottom"
    
)
// LayoutColoraxisColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type LayoutColoraxisColorbarXanchor string 

const (
    LayoutColoraxisColorbarXanchorLeft LayoutColoraxisColorbarXanchor = "left"
    LayoutColoraxisColorbarXanchorCenter LayoutColoraxisColorbarXanchor = "center"
    LayoutColoraxisColorbarXanchorRight LayoutColoraxisColorbarXanchor = "right"
    
)
// LayoutColoraxisColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type LayoutColoraxisColorbarYanchor string 

const (
    LayoutColoraxisColorbarYanchorTop LayoutColoraxisColorbarYanchor = "top"
    LayoutColoraxisColorbarYanchorMiddle LayoutColoraxisColorbarYanchor = "middle"
    LayoutColoraxisColorbarYanchorBottom LayoutColoraxisColorbarYanchor = "bottom"
    
)
// LayoutDirection Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the direction corresponding to positive angles in legacy polar charts.
type LayoutDirection string 

const (
    LayoutDirectionClockwise LayoutDirection = "clockwise"
    LayoutDirectionCounterclockwise LayoutDirection = "counterclockwise"
    
)
// LayoutDragmode Determines the mode of drag interactions. *select* and *lasso* apply only to scatter traces with markers or text. *orbit* and *turntable* apply only to 3D scenes.
type LayoutDragmode interface{} 

var (
    LayoutDragmodeZoom LayoutDragmode = "zoom"
    LayoutDragmodePan LayoutDragmode = "pan"
    LayoutDragmodeSelect LayoutDragmode = "select"
    LayoutDragmodeLasso LayoutDragmode = "lasso"
    LayoutDragmodeDrawclosedpath LayoutDragmode = "drawclosedpath"
    LayoutDragmodeDrawopenpath LayoutDragmode = "drawopenpath"
    LayoutDragmodeDrawline LayoutDragmode = "drawline"
    LayoutDragmodeDrawrect LayoutDragmode = "drawrect"
    LayoutDragmodeDrawcircle LayoutDragmode = "drawcircle"
    LayoutDragmodeOrbit LayoutDragmode = "orbit"
    LayoutDragmodeTurntable LayoutDragmode = "turntable"
    LayoutDragmodeFalse LayoutDragmode = false
    
)
// LayoutGeoFitbounds Determines if this subplot's view settings are auto-computed to fit trace data. On scoped maps, setting `fitbounds` leads to `center.lon` and `center.lat` getting auto-filled. On maps with a non-clipped projection, setting `fitbounds` leads to `center.lon`, `center.lat`, and `projection.rotation.lon` getting auto-filled. On maps with a clipped projection, setting `fitbounds` leads to `center.lon`, `center.lat`, `projection.rotation.lon`, `projection.rotation.lat`, `lonaxis.range` and `lonaxis.range` getting auto-filled. If *locations*, only the trace's visible locations are considered in the `fitbounds` computations. If *geojson*, the entire trace input `geojson` (if provided) is considered in the `fitbounds` computations, Defaults to *false*.
type LayoutGeoFitbounds interface{} 

var (
    LayoutGeoFitboundsFalse LayoutGeoFitbounds = false
    LayoutGeoFitboundsLocations LayoutGeoFitbounds = "locations"
    LayoutGeoFitboundsGeojson LayoutGeoFitbounds = "geojson"
    
)
// LayoutGeoProjectionType Sets the projection type.
type LayoutGeoProjectionType string 

const (
    LayoutGeoProjectionTypeEquirectangular LayoutGeoProjectionType = "equirectangular"
    LayoutGeoProjectionTypeMercator LayoutGeoProjectionType = "mercator"
    LayoutGeoProjectionTypeOrthographic LayoutGeoProjectionType = "orthographic"
    LayoutGeoProjectionTypeNaturalEarth LayoutGeoProjectionType = "natural earth"
    LayoutGeoProjectionTypeKavrayskiy7 LayoutGeoProjectionType = "kavrayskiy7"
    LayoutGeoProjectionTypeMiller LayoutGeoProjectionType = "miller"
    LayoutGeoProjectionTypeRobinson LayoutGeoProjectionType = "robinson"
    LayoutGeoProjectionTypeEckert4 LayoutGeoProjectionType = "eckert4"
    LayoutGeoProjectionTypeAzimuthalEqualArea LayoutGeoProjectionType = "azimuthal equal area"
    LayoutGeoProjectionTypeAzimuthalEquidistant LayoutGeoProjectionType = "azimuthal equidistant"
    LayoutGeoProjectionTypeConicEqualArea LayoutGeoProjectionType = "conic equal area"
    LayoutGeoProjectionTypeConicConformal LayoutGeoProjectionType = "conic conformal"
    LayoutGeoProjectionTypeConicEquidistant LayoutGeoProjectionType = "conic equidistant"
    LayoutGeoProjectionTypeGnomonic LayoutGeoProjectionType = "gnomonic"
    LayoutGeoProjectionTypeStereographic LayoutGeoProjectionType = "stereographic"
    LayoutGeoProjectionTypeMollweide LayoutGeoProjectionType = "mollweide"
    LayoutGeoProjectionTypeHammer LayoutGeoProjectionType = "hammer"
    LayoutGeoProjectionTypeTransverseMercator LayoutGeoProjectionType = "transverse mercator"
    LayoutGeoProjectionTypeAlbersUsa LayoutGeoProjectionType = "albers usa"
    LayoutGeoProjectionTypeWinkelTripel LayoutGeoProjectionType = "winkel tripel"
    LayoutGeoProjectionTypeAitoff LayoutGeoProjectionType = "aitoff"
    LayoutGeoProjectionTypeSinusoidal LayoutGeoProjectionType = "sinusoidal"
    
)
// LayoutGeoResolution Sets the resolution of the base layers. The values have units of km/mm e.g. 110 corresponds to a scale ratio of 1:110,000,000.
type LayoutGeoResolution interface{} 

var (
    LayoutGeoResolutionNumber110 LayoutGeoResolution = 110
    LayoutGeoResolutionNumber50 LayoutGeoResolution = 50
    
)
// LayoutGeoScope Set the scope of the map.
type LayoutGeoScope string 

const (
    LayoutGeoScopeWorld LayoutGeoScope = "world"
    LayoutGeoScopeUsa LayoutGeoScope = "usa"
    LayoutGeoScopeEurope LayoutGeoScope = "europe"
    LayoutGeoScopeAsia LayoutGeoScope = "asia"
    LayoutGeoScopeAfrica LayoutGeoScope = "africa"
    LayoutGeoScopeNorthAmerica LayoutGeoScope = "north america"
    LayoutGeoScopeSouthAmerica LayoutGeoScope = "south america"
    
)
// LayoutGridPattern If no `subplots`, `xaxes`, or `yaxes` are given but we do have `rows` and `columns`, we can generate defaults using consecutive axis IDs, in two ways: *coupled* gives one x axis per column and one y axis per row. *independent* uses a new xy pair for each cell, left-to-right across each row then iterating rows according to `roworder`.
type LayoutGridPattern string 

const (
    LayoutGridPatternIndependent LayoutGridPattern = "independent"
    LayoutGridPatternCoupled LayoutGridPattern = "coupled"
    
)
// LayoutGridRoworder Is the first row the top or the bottom? Note that columns are always enumerated from left to right.
type LayoutGridRoworder string 

const (
    LayoutGridRoworderTopToBottom LayoutGridRoworder = "top to bottom"
    LayoutGridRoworderBottomToTop LayoutGridRoworder = "bottom to top"
    
)
// LayoutGridXside Sets where the x axis labels and titles go. *bottom* means the very bottom of the grid. *bottom plot* is the lowest plot that each x axis is used in. *top* and *top plot* are similar.
type LayoutGridXside string 

const (
    LayoutGridXsideBottom LayoutGridXside = "bottom"
    LayoutGridXsideBottomPlot LayoutGridXside = "bottom plot"
    LayoutGridXsideTopPlot LayoutGridXside = "top plot"
    LayoutGridXsideTop LayoutGridXside = "top"
    
)
// LayoutGridYside Sets where the y axis labels and titles go. *left* means the very left edge of the grid. *left plot* is the leftmost plot that each y axis is used in. *right* and *right plot* are similar.
type LayoutGridYside string 

const (
    LayoutGridYsideLeft LayoutGridYside = "left"
    LayoutGridYsideLeftPlot LayoutGridYside = "left plot"
    LayoutGridYsideRightPlot LayoutGridYside = "right plot"
    LayoutGridYsideRight LayoutGridYside = "right"
    
)
// LayoutHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type LayoutHoverlabelAlign string 

const (
    LayoutHoverlabelAlignLeft LayoutHoverlabelAlign = "left"
    LayoutHoverlabelAlignRight LayoutHoverlabelAlign = "right"
    LayoutHoverlabelAlignAuto LayoutHoverlabelAlign = "auto"
    
)
// LayoutHovermode Determines the mode of hover interactions. If *closest*, a single hoverlabel will appear for the *closest* point within the `hoverdistance`. If *x* (or *y*), multiple hoverlabels will appear for multiple points at the *closest* x- (or y-) coordinate within the `hoverdistance`, with the caveat that no more than one hoverlabel will appear per trace. If *x unified* (or *y unified*), a single hoverlabel will appear multiple points at the closest x- (or y-) coordinate within the `hoverdistance` with the caveat that no more than one hoverlabel will appear per trace. In this mode, spikelines are enabled by default perpendicular to the specified axis. If false, hover interactions are disabled. If `clickmode` includes the *select* flag, `hovermode` defaults to *closest*. If `clickmode` lacks the *select* flag, it defaults to *x* or *y* (depending on the trace's `orientation` value) for plots based on cartesian coordinates. For anything else the default value is *closest*.
type LayoutHovermode interface{} 

var (
    LayoutHovermodeX LayoutHovermode = "x"
    LayoutHovermodeY LayoutHovermode = "y"
    LayoutHovermodeClosest LayoutHovermode = "closest"
    LayoutHovermodeFalse LayoutHovermode = false
    LayoutHovermodeXUnified LayoutHovermode = "x unified"
    LayoutHovermodeYUnified LayoutHovermode = "y unified"
    
)
// LayoutLegendItemclick Determines the behavior on legend item click. *toggle* toggles the visibility of the item clicked on the graph. *toggleothers* makes the clicked item the sole visible item on the graph. *false* disable legend item click interactions.
type LayoutLegendItemclick interface{} 

var (
    LayoutLegendItemclickToggle LayoutLegendItemclick = "toggle"
    LayoutLegendItemclickToggleothers LayoutLegendItemclick = "toggleothers"
    LayoutLegendItemclickFalse LayoutLegendItemclick = false
    
)
// LayoutLegendItemdoubleclick Determines the behavior on legend item double-click. *toggle* toggles the visibility of the item clicked on the graph. *toggleothers* makes the clicked item the sole visible item on the graph. *false* disable legend item double-click interactions.
type LayoutLegendItemdoubleclick interface{} 

var (
    LayoutLegendItemdoubleclickToggle LayoutLegendItemdoubleclick = "toggle"
    LayoutLegendItemdoubleclickToggleothers LayoutLegendItemdoubleclick = "toggleothers"
    LayoutLegendItemdoubleclickFalse LayoutLegendItemdoubleclick = false
    
)
// LayoutLegendItemsizing Determines if the legend items symbols scale with their corresponding *trace* attributes or remain *constant* independent of the symbol size on the graph.
type LayoutLegendItemsizing string 

const (
    LayoutLegendItemsizingTrace LayoutLegendItemsizing = "trace"
    LayoutLegendItemsizingConstant LayoutLegendItemsizing = "constant"
    
)
// LayoutLegendOrientation Sets the orientation of the legend.
type LayoutLegendOrientation string 

const (
    LayoutLegendOrientationV LayoutLegendOrientation = "v"
    LayoutLegendOrientationH LayoutLegendOrientation = "h"
    
)
// LayoutLegendTitleSide Determines the location of legend's title with respect to the legend items. Defaulted to *top* with `orientation` is *h*. Defaulted to *left* with `orientation` is *v*. The *top left* options could be used to expand legend area in both x and y sides.
type LayoutLegendTitleSide string 

const (
    LayoutLegendTitleSideTop LayoutLegendTitleSide = "top"
    LayoutLegendTitleSideLeft LayoutLegendTitleSide = "left"
    LayoutLegendTitleSideTopLeft LayoutLegendTitleSide = "top left"
    
)
// LayoutLegendValign Sets the vertical alignment of the symbols with respect to their associated text.
type LayoutLegendValign string 

const (
    LayoutLegendValignTop LayoutLegendValign = "top"
    LayoutLegendValignMiddle LayoutLegendValign = "middle"
    LayoutLegendValignBottom LayoutLegendValign = "bottom"
    
)
// LayoutLegendXanchor Sets the legend's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the legend. Value *auto* anchors legends to the right for `x` values greater than or equal to 2/3, anchors legends to the left for `x` values less than or equal to 1/3 and anchors legends with respect to their center otherwise.
type LayoutLegendXanchor string 

const (
    LayoutLegendXanchorAuto LayoutLegendXanchor = "auto"
    LayoutLegendXanchorLeft LayoutLegendXanchor = "left"
    LayoutLegendXanchorCenter LayoutLegendXanchor = "center"
    LayoutLegendXanchorRight LayoutLegendXanchor = "right"
    
)
// LayoutLegendYanchor Sets the legend's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the legend. Value *auto* anchors legends at their bottom for `y` values less than or equal to 1/3, anchors legends to at their top for `y` values greater than or equal to 2/3 and anchors legends with respect to their middle otherwise.
type LayoutLegendYanchor string 

const (
    LayoutLegendYanchorAuto LayoutLegendYanchor = "auto"
    LayoutLegendYanchorTop LayoutLegendYanchor = "top"
    LayoutLegendYanchorMiddle LayoutLegendYanchor = "middle"
    LayoutLegendYanchorBottom LayoutLegendYanchor = "bottom"
    
)
// LayoutModebarOrientation Sets the orientation of the modebar.
type LayoutModebarOrientation string 

const (
    LayoutModebarOrientationV LayoutModebarOrientation = "v"
    LayoutModebarOrientationH LayoutModebarOrientation = "h"
    
)
// LayoutNewshapeDrawdirection When `dragmode` is set to *drawrect*, *drawline* or *drawcircle* this limits the drag to be horizontal, vertical or diagonal. Using *diagonal* there is no limit e.g. in drawing lines in any direction. *ortho* limits the draw to be either horizontal or vertical. *horizontal* allows horizontal extend. *vertical* allows vertical extend.
type LayoutNewshapeDrawdirection string 

const (
    LayoutNewshapeDrawdirectionOrtho LayoutNewshapeDrawdirection = "ortho"
    LayoutNewshapeDrawdirectionHorizontal LayoutNewshapeDrawdirection = "horizontal"
    LayoutNewshapeDrawdirectionVertical LayoutNewshapeDrawdirection = "vertical"
    LayoutNewshapeDrawdirectionDiagonal LayoutNewshapeDrawdirection = "diagonal"
    
)
// LayoutNewshapeFillrule Determines the path's interior. For more info please visit https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill-rule
type LayoutNewshapeFillrule string 

const (
    LayoutNewshapeFillruleEvenodd LayoutNewshapeFillrule = "evenodd"
    LayoutNewshapeFillruleNonzero LayoutNewshapeFillrule = "nonzero"
    
)
// LayoutNewshapeLayer Specifies whether new shapes are drawn below or above traces.
type LayoutNewshapeLayer string 

const (
    LayoutNewshapeLayerBelow LayoutNewshapeLayer = "below"
    LayoutNewshapeLayerAbove LayoutNewshapeLayer = "above"
    
)
// LayoutPolarAngularaxisAutotypenumbers Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
type LayoutPolarAngularaxisAutotypenumbers string 

const (
    LayoutPolarAngularaxisAutotypenumbersConvertTypes LayoutPolarAngularaxisAutotypenumbers = "convert types"
    LayoutPolarAngularaxisAutotypenumbersStrict LayoutPolarAngularaxisAutotypenumbers = "strict"
    
)
// LayoutPolarAngularaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutPolarAngularaxisCategoryorder string 

const (
    LayoutPolarAngularaxisCategoryorderTrace LayoutPolarAngularaxisCategoryorder = "trace"
    LayoutPolarAngularaxisCategoryorderCategoryAscending LayoutPolarAngularaxisCategoryorder = "category ascending"
    LayoutPolarAngularaxisCategoryorderCategoryDescending LayoutPolarAngularaxisCategoryorder = "category descending"
    LayoutPolarAngularaxisCategoryorderArray LayoutPolarAngularaxisCategoryorder = "array"
    LayoutPolarAngularaxisCategoryorderTotalAscending LayoutPolarAngularaxisCategoryorder = "total ascending"
    LayoutPolarAngularaxisCategoryorderTotalDescending LayoutPolarAngularaxisCategoryorder = "total descending"
    LayoutPolarAngularaxisCategoryorderMinAscending LayoutPolarAngularaxisCategoryorder = "min ascending"
    LayoutPolarAngularaxisCategoryorderMinDescending LayoutPolarAngularaxisCategoryorder = "min descending"
    LayoutPolarAngularaxisCategoryorderMaxAscending LayoutPolarAngularaxisCategoryorder = "max ascending"
    LayoutPolarAngularaxisCategoryorderMaxDescending LayoutPolarAngularaxisCategoryorder = "max descending"
    LayoutPolarAngularaxisCategoryorderSumAscending LayoutPolarAngularaxisCategoryorder = "sum ascending"
    LayoutPolarAngularaxisCategoryorderSumDescending LayoutPolarAngularaxisCategoryorder = "sum descending"
    LayoutPolarAngularaxisCategoryorderMeanAscending LayoutPolarAngularaxisCategoryorder = "mean ascending"
    LayoutPolarAngularaxisCategoryorderMeanDescending LayoutPolarAngularaxisCategoryorder = "mean descending"
    LayoutPolarAngularaxisCategoryorderMedianAscending LayoutPolarAngularaxisCategoryorder = "median ascending"
    LayoutPolarAngularaxisCategoryorderMedianDescending LayoutPolarAngularaxisCategoryorder = "median descending"
    
)
// LayoutPolarAngularaxisDirection Sets the direction corresponding to positive angles.
type LayoutPolarAngularaxisDirection string 

const (
    LayoutPolarAngularaxisDirectionCounterclockwise LayoutPolarAngularaxisDirection = "counterclockwise"
    LayoutPolarAngularaxisDirectionClockwise LayoutPolarAngularaxisDirection = "clockwise"
    
)
// LayoutPolarAngularaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutPolarAngularaxisExponentformat string 

const (
    LayoutPolarAngularaxisExponentformatNone LayoutPolarAngularaxisExponentformat = "none"
    LayoutPolarAngularaxisExponentformatE1 LayoutPolarAngularaxisExponentformat = "e"
    LayoutPolarAngularaxisExponentformatE2 LayoutPolarAngularaxisExponentformat = "E"
    LayoutPolarAngularaxisExponentformatPower LayoutPolarAngularaxisExponentformat = "power"
    LayoutPolarAngularaxisExponentformatSi LayoutPolarAngularaxisExponentformat = "SI"
    LayoutPolarAngularaxisExponentformatB LayoutPolarAngularaxisExponentformat = "B"
    
)
// LayoutPolarAngularaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutPolarAngularaxisLayer string 

const (
    LayoutPolarAngularaxisLayerAboveTraces LayoutPolarAngularaxisLayer = "above traces"
    LayoutPolarAngularaxisLayerBelowTraces LayoutPolarAngularaxisLayer = "below traces"
    
)
// LayoutPolarAngularaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutPolarAngularaxisShowexponent string 

const (
    LayoutPolarAngularaxisShowexponentAll LayoutPolarAngularaxisShowexponent = "all"
    LayoutPolarAngularaxisShowexponentFirst LayoutPolarAngularaxisShowexponent = "first"
    LayoutPolarAngularaxisShowexponentLast LayoutPolarAngularaxisShowexponent = "last"
    LayoutPolarAngularaxisShowexponentNone LayoutPolarAngularaxisShowexponent = "none"
    
)
// LayoutPolarAngularaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutPolarAngularaxisShowtickprefix string 

const (
    LayoutPolarAngularaxisShowtickprefixAll LayoutPolarAngularaxisShowtickprefix = "all"
    LayoutPolarAngularaxisShowtickprefixFirst LayoutPolarAngularaxisShowtickprefix = "first"
    LayoutPolarAngularaxisShowtickprefixLast LayoutPolarAngularaxisShowtickprefix = "last"
    LayoutPolarAngularaxisShowtickprefixNone LayoutPolarAngularaxisShowtickprefix = "none"
    
)
// LayoutPolarAngularaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutPolarAngularaxisShowticksuffix string 

const (
    LayoutPolarAngularaxisShowticksuffixAll LayoutPolarAngularaxisShowticksuffix = "all"
    LayoutPolarAngularaxisShowticksuffixFirst LayoutPolarAngularaxisShowticksuffix = "first"
    LayoutPolarAngularaxisShowticksuffixLast LayoutPolarAngularaxisShowticksuffix = "last"
    LayoutPolarAngularaxisShowticksuffixNone LayoutPolarAngularaxisShowticksuffix = "none"
    
)
// LayoutPolarAngularaxisThetaunit Sets the format unit of the formatted *theta* values. Has an effect only when `angularaxis.type` is *linear*.
type LayoutPolarAngularaxisThetaunit string 

const (
    LayoutPolarAngularaxisThetaunitRadians LayoutPolarAngularaxisThetaunit = "radians"
    LayoutPolarAngularaxisThetaunitDegrees LayoutPolarAngularaxisThetaunit = "degrees"
    
)
// LayoutPolarAngularaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutPolarAngularaxisTickmode string 

const (
    LayoutPolarAngularaxisTickmodeAuto LayoutPolarAngularaxisTickmode = "auto"
    LayoutPolarAngularaxisTickmodeLinear LayoutPolarAngularaxisTickmode = "linear"
    LayoutPolarAngularaxisTickmodeArray LayoutPolarAngularaxisTickmode = "array"
    
)
// LayoutPolarAngularaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutPolarAngularaxisTicks string 

const (
    LayoutPolarAngularaxisTicksOutside LayoutPolarAngularaxisTicks = "outside"
    LayoutPolarAngularaxisTicksInside LayoutPolarAngularaxisTicks = "inside"
    LayoutPolarAngularaxisTicksEmpty LayoutPolarAngularaxisTicks = ""
    
)
// LayoutPolarAngularaxisType Sets the angular axis type. If *linear*, set `thetaunit` to determine the unit in which axis value are shown. If *category, use `period` to set the number of integer coordinates around polar axis.
type LayoutPolarAngularaxisType string 

const (
    LayoutPolarAngularaxisTypeHyphenHyphen LayoutPolarAngularaxisType = "-"
    LayoutPolarAngularaxisTypeLinear LayoutPolarAngularaxisType = "linear"
    LayoutPolarAngularaxisTypeCategory LayoutPolarAngularaxisType = "category"
    
)
// LayoutPolarGridshape Determines if the radial axis grid lines and angular axis line are drawn as *circular* sectors or as *linear* (polygon) sectors. Has an effect only when the angular axis has `type` *category*. Note that `radialaxis.angle` is snapped to the angle of the closest vertex when `gridshape` is *circular* (so that radial axis scale is the same as the data scale).
type LayoutPolarGridshape string 

const (
    LayoutPolarGridshapeCircular LayoutPolarGridshape = "circular"
    LayoutPolarGridshapeLinear LayoutPolarGridshape = "linear"
    
)
// LayoutPolarRadialaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutPolarRadialaxisAutorange interface{} 

var (
    LayoutPolarRadialaxisAutorangeTrue LayoutPolarRadialaxisAutorange = true
    LayoutPolarRadialaxisAutorangeFalse LayoutPolarRadialaxisAutorange = false
    LayoutPolarRadialaxisAutorangeReversed LayoutPolarRadialaxisAutorange = "reversed"
    
)
// LayoutPolarRadialaxisAutotypenumbers Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
type LayoutPolarRadialaxisAutotypenumbers string 

const (
    LayoutPolarRadialaxisAutotypenumbersConvertTypes LayoutPolarRadialaxisAutotypenumbers = "convert types"
    LayoutPolarRadialaxisAutotypenumbersStrict LayoutPolarRadialaxisAutotypenumbers = "strict"
    
)
// LayoutPolarRadialaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutPolarRadialaxisCalendar string 

const (
    LayoutPolarRadialaxisCalendarGregorian LayoutPolarRadialaxisCalendar = "gregorian"
    LayoutPolarRadialaxisCalendarChinese LayoutPolarRadialaxisCalendar = "chinese"
    LayoutPolarRadialaxisCalendarCoptic LayoutPolarRadialaxisCalendar = "coptic"
    LayoutPolarRadialaxisCalendarDiscworld LayoutPolarRadialaxisCalendar = "discworld"
    LayoutPolarRadialaxisCalendarEthiopian LayoutPolarRadialaxisCalendar = "ethiopian"
    LayoutPolarRadialaxisCalendarHebrew LayoutPolarRadialaxisCalendar = "hebrew"
    LayoutPolarRadialaxisCalendarIslamic LayoutPolarRadialaxisCalendar = "islamic"
    LayoutPolarRadialaxisCalendarJulian LayoutPolarRadialaxisCalendar = "julian"
    LayoutPolarRadialaxisCalendarMayan LayoutPolarRadialaxisCalendar = "mayan"
    LayoutPolarRadialaxisCalendarNanakshahi LayoutPolarRadialaxisCalendar = "nanakshahi"
    LayoutPolarRadialaxisCalendarNepali LayoutPolarRadialaxisCalendar = "nepali"
    LayoutPolarRadialaxisCalendarPersian LayoutPolarRadialaxisCalendar = "persian"
    LayoutPolarRadialaxisCalendarJalali LayoutPolarRadialaxisCalendar = "jalali"
    LayoutPolarRadialaxisCalendarTaiwan LayoutPolarRadialaxisCalendar = "taiwan"
    LayoutPolarRadialaxisCalendarThai LayoutPolarRadialaxisCalendar = "thai"
    LayoutPolarRadialaxisCalendarUmmalqura LayoutPolarRadialaxisCalendar = "ummalqura"
    
)
// LayoutPolarRadialaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutPolarRadialaxisCategoryorder string 

const (
    LayoutPolarRadialaxisCategoryorderTrace LayoutPolarRadialaxisCategoryorder = "trace"
    LayoutPolarRadialaxisCategoryorderCategoryAscending LayoutPolarRadialaxisCategoryorder = "category ascending"
    LayoutPolarRadialaxisCategoryorderCategoryDescending LayoutPolarRadialaxisCategoryorder = "category descending"
    LayoutPolarRadialaxisCategoryorderArray LayoutPolarRadialaxisCategoryorder = "array"
    LayoutPolarRadialaxisCategoryorderTotalAscending LayoutPolarRadialaxisCategoryorder = "total ascending"
    LayoutPolarRadialaxisCategoryorderTotalDescending LayoutPolarRadialaxisCategoryorder = "total descending"
    LayoutPolarRadialaxisCategoryorderMinAscending LayoutPolarRadialaxisCategoryorder = "min ascending"
    LayoutPolarRadialaxisCategoryorderMinDescending LayoutPolarRadialaxisCategoryorder = "min descending"
    LayoutPolarRadialaxisCategoryorderMaxAscending LayoutPolarRadialaxisCategoryorder = "max ascending"
    LayoutPolarRadialaxisCategoryorderMaxDescending LayoutPolarRadialaxisCategoryorder = "max descending"
    LayoutPolarRadialaxisCategoryorderSumAscending LayoutPolarRadialaxisCategoryorder = "sum ascending"
    LayoutPolarRadialaxisCategoryorderSumDescending LayoutPolarRadialaxisCategoryorder = "sum descending"
    LayoutPolarRadialaxisCategoryorderMeanAscending LayoutPolarRadialaxisCategoryorder = "mean ascending"
    LayoutPolarRadialaxisCategoryorderMeanDescending LayoutPolarRadialaxisCategoryorder = "mean descending"
    LayoutPolarRadialaxisCategoryorderMedianAscending LayoutPolarRadialaxisCategoryorder = "median ascending"
    LayoutPolarRadialaxisCategoryorderMedianDescending LayoutPolarRadialaxisCategoryorder = "median descending"
    
)
// LayoutPolarRadialaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutPolarRadialaxisExponentformat string 

const (
    LayoutPolarRadialaxisExponentformatNone LayoutPolarRadialaxisExponentformat = "none"
    LayoutPolarRadialaxisExponentformatE1 LayoutPolarRadialaxisExponentformat = "e"
    LayoutPolarRadialaxisExponentformatE2 LayoutPolarRadialaxisExponentformat = "E"
    LayoutPolarRadialaxisExponentformatPower LayoutPolarRadialaxisExponentformat = "power"
    LayoutPolarRadialaxisExponentformatSi LayoutPolarRadialaxisExponentformat = "SI"
    LayoutPolarRadialaxisExponentformatB LayoutPolarRadialaxisExponentformat = "B"
    
)
// LayoutPolarRadialaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutPolarRadialaxisLayer string 

const (
    LayoutPolarRadialaxisLayerAboveTraces LayoutPolarRadialaxisLayer = "above traces"
    LayoutPolarRadialaxisLayerBelowTraces LayoutPolarRadialaxisLayer = "below traces"
    
)
// LayoutPolarRadialaxisRangemode If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. If *normal*, the range is computed in relation to the extrema of the input data (same behavior as for cartesian axes).
type LayoutPolarRadialaxisRangemode string 

const (
    LayoutPolarRadialaxisRangemodeTozero LayoutPolarRadialaxisRangemode = "tozero"
    LayoutPolarRadialaxisRangemodeNonnegative LayoutPolarRadialaxisRangemode = "nonnegative"
    LayoutPolarRadialaxisRangemodeNormal LayoutPolarRadialaxisRangemode = "normal"
    
)
// LayoutPolarRadialaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutPolarRadialaxisShowexponent string 

const (
    LayoutPolarRadialaxisShowexponentAll LayoutPolarRadialaxisShowexponent = "all"
    LayoutPolarRadialaxisShowexponentFirst LayoutPolarRadialaxisShowexponent = "first"
    LayoutPolarRadialaxisShowexponentLast LayoutPolarRadialaxisShowexponent = "last"
    LayoutPolarRadialaxisShowexponentNone LayoutPolarRadialaxisShowexponent = "none"
    
)
// LayoutPolarRadialaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutPolarRadialaxisShowtickprefix string 

const (
    LayoutPolarRadialaxisShowtickprefixAll LayoutPolarRadialaxisShowtickprefix = "all"
    LayoutPolarRadialaxisShowtickprefixFirst LayoutPolarRadialaxisShowtickprefix = "first"
    LayoutPolarRadialaxisShowtickprefixLast LayoutPolarRadialaxisShowtickprefix = "last"
    LayoutPolarRadialaxisShowtickprefixNone LayoutPolarRadialaxisShowtickprefix = "none"
    
)
// LayoutPolarRadialaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutPolarRadialaxisShowticksuffix string 

const (
    LayoutPolarRadialaxisShowticksuffixAll LayoutPolarRadialaxisShowticksuffix = "all"
    LayoutPolarRadialaxisShowticksuffixFirst LayoutPolarRadialaxisShowticksuffix = "first"
    LayoutPolarRadialaxisShowticksuffixLast LayoutPolarRadialaxisShowticksuffix = "last"
    LayoutPolarRadialaxisShowticksuffixNone LayoutPolarRadialaxisShowticksuffix = "none"
    
)
// LayoutPolarRadialaxisSide Determines on which side of radial axis line the tick and tick labels appear.
type LayoutPolarRadialaxisSide string 

const (
    LayoutPolarRadialaxisSideClockwise LayoutPolarRadialaxisSide = "clockwise"
    LayoutPolarRadialaxisSideCounterclockwise LayoutPolarRadialaxisSide = "counterclockwise"
    
)
// LayoutPolarRadialaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutPolarRadialaxisTickmode string 

const (
    LayoutPolarRadialaxisTickmodeAuto LayoutPolarRadialaxisTickmode = "auto"
    LayoutPolarRadialaxisTickmodeLinear LayoutPolarRadialaxisTickmode = "linear"
    LayoutPolarRadialaxisTickmodeArray LayoutPolarRadialaxisTickmode = "array"
    
)
// LayoutPolarRadialaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutPolarRadialaxisTicks string 

const (
    LayoutPolarRadialaxisTicksOutside LayoutPolarRadialaxisTicks = "outside"
    LayoutPolarRadialaxisTicksInside LayoutPolarRadialaxisTicks = "inside"
    LayoutPolarRadialaxisTicksEmpty LayoutPolarRadialaxisTicks = ""
    
)
// LayoutPolarRadialaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutPolarRadialaxisType string 

const (
    LayoutPolarRadialaxisTypeHyphenHyphen LayoutPolarRadialaxisType = "-"
    LayoutPolarRadialaxisTypeLinear LayoutPolarRadialaxisType = "linear"
    LayoutPolarRadialaxisTypeLog LayoutPolarRadialaxisType = "log"
    LayoutPolarRadialaxisTypeDate LayoutPolarRadialaxisType = "date"
    LayoutPolarRadialaxisTypeCategory LayoutPolarRadialaxisType = "category"
    
)
// LayoutRadialaxisTickorientation Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the orientation (from the paper perspective) of the radial axis tick labels.
type LayoutRadialaxisTickorientation string 

const (
    LayoutRadialaxisTickorientationHorizontal LayoutRadialaxisTickorientation = "horizontal"
    LayoutRadialaxisTickorientationVertical LayoutRadialaxisTickorientation = "vertical"
    
)
// LayoutSceneAspectmode If *cube*, this scene's axes are drawn as a cube, regardless of the axes' ranges. If *data*, this scene's axes are drawn in proportion with the axes' ranges. If *manual*, this scene's axes are drawn in proportion with the input of *aspectratio* (the default behavior if *aspectratio* is provided). If *auto*, this scene's axes are drawn using the results of *data* except when one axis is more than four times the size of the two others, where in that case the results of *cube* are used.
type LayoutSceneAspectmode string 

const (
    LayoutSceneAspectmodeAuto LayoutSceneAspectmode = "auto"
    LayoutSceneAspectmodeCube LayoutSceneAspectmode = "cube"
    LayoutSceneAspectmodeData LayoutSceneAspectmode = "data"
    LayoutSceneAspectmodeManual LayoutSceneAspectmode = "manual"
    
)
// LayoutSceneCameraProjectionType Sets the projection type. The projection type could be either *perspective* or *orthographic*. The default is *perspective*.
type LayoutSceneCameraProjectionType string 

const (
    LayoutSceneCameraProjectionTypePerspective LayoutSceneCameraProjectionType = "perspective"
    LayoutSceneCameraProjectionTypeOrthographic LayoutSceneCameraProjectionType = "orthographic"
    
)
// LayoutSceneDragmode Determines the mode of drag interactions for this scene.
type LayoutSceneDragmode interface{} 

var (
    LayoutSceneDragmodeOrbit LayoutSceneDragmode = "orbit"
    LayoutSceneDragmodeTurntable LayoutSceneDragmode = "turntable"
    LayoutSceneDragmodeZoom LayoutSceneDragmode = "zoom"
    LayoutSceneDragmodePan LayoutSceneDragmode = "pan"
    LayoutSceneDragmodeFalse LayoutSceneDragmode = false
    
)
// LayoutSceneHovermode Determines the mode of hover interactions for this scene.
type LayoutSceneHovermode interface{} 

var (
    LayoutSceneHovermodeClosest LayoutSceneHovermode = "closest"
    LayoutSceneHovermodeFalse LayoutSceneHovermode = false
    
)
// LayoutSceneXaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutSceneXaxisAutorange interface{} 

var (
    LayoutSceneXaxisAutorangeTrue LayoutSceneXaxisAutorange = true
    LayoutSceneXaxisAutorangeFalse LayoutSceneXaxisAutorange = false
    LayoutSceneXaxisAutorangeReversed LayoutSceneXaxisAutorange = "reversed"
    
)
// LayoutSceneXaxisAutotypenumbers Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
type LayoutSceneXaxisAutotypenumbers string 

const (
    LayoutSceneXaxisAutotypenumbersConvertTypes LayoutSceneXaxisAutotypenumbers = "convert types"
    LayoutSceneXaxisAutotypenumbersStrict LayoutSceneXaxisAutotypenumbers = "strict"
    
)
// LayoutSceneXaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutSceneXaxisCalendar string 

const (
    LayoutSceneXaxisCalendarGregorian LayoutSceneXaxisCalendar = "gregorian"
    LayoutSceneXaxisCalendarChinese LayoutSceneXaxisCalendar = "chinese"
    LayoutSceneXaxisCalendarCoptic LayoutSceneXaxisCalendar = "coptic"
    LayoutSceneXaxisCalendarDiscworld LayoutSceneXaxisCalendar = "discworld"
    LayoutSceneXaxisCalendarEthiopian LayoutSceneXaxisCalendar = "ethiopian"
    LayoutSceneXaxisCalendarHebrew LayoutSceneXaxisCalendar = "hebrew"
    LayoutSceneXaxisCalendarIslamic LayoutSceneXaxisCalendar = "islamic"
    LayoutSceneXaxisCalendarJulian LayoutSceneXaxisCalendar = "julian"
    LayoutSceneXaxisCalendarMayan LayoutSceneXaxisCalendar = "mayan"
    LayoutSceneXaxisCalendarNanakshahi LayoutSceneXaxisCalendar = "nanakshahi"
    LayoutSceneXaxisCalendarNepali LayoutSceneXaxisCalendar = "nepali"
    LayoutSceneXaxisCalendarPersian LayoutSceneXaxisCalendar = "persian"
    LayoutSceneXaxisCalendarJalali LayoutSceneXaxisCalendar = "jalali"
    LayoutSceneXaxisCalendarTaiwan LayoutSceneXaxisCalendar = "taiwan"
    LayoutSceneXaxisCalendarThai LayoutSceneXaxisCalendar = "thai"
    LayoutSceneXaxisCalendarUmmalqura LayoutSceneXaxisCalendar = "ummalqura"
    
)
// LayoutSceneXaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutSceneXaxisCategoryorder string 

const (
    LayoutSceneXaxisCategoryorderTrace LayoutSceneXaxisCategoryorder = "trace"
    LayoutSceneXaxisCategoryorderCategoryAscending LayoutSceneXaxisCategoryorder = "category ascending"
    LayoutSceneXaxisCategoryorderCategoryDescending LayoutSceneXaxisCategoryorder = "category descending"
    LayoutSceneXaxisCategoryorderArray LayoutSceneXaxisCategoryorder = "array"
    LayoutSceneXaxisCategoryorderTotalAscending LayoutSceneXaxisCategoryorder = "total ascending"
    LayoutSceneXaxisCategoryorderTotalDescending LayoutSceneXaxisCategoryorder = "total descending"
    LayoutSceneXaxisCategoryorderMinAscending LayoutSceneXaxisCategoryorder = "min ascending"
    LayoutSceneXaxisCategoryorderMinDescending LayoutSceneXaxisCategoryorder = "min descending"
    LayoutSceneXaxisCategoryorderMaxAscending LayoutSceneXaxisCategoryorder = "max ascending"
    LayoutSceneXaxisCategoryorderMaxDescending LayoutSceneXaxisCategoryorder = "max descending"
    LayoutSceneXaxisCategoryorderSumAscending LayoutSceneXaxisCategoryorder = "sum ascending"
    LayoutSceneXaxisCategoryorderSumDescending LayoutSceneXaxisCategoryorder = "sum descending"
    LayoutSceneXaxisCategoryorderMeanAscending LayoutSceneXaxisCategoryorder = "mean ascending"
    LayoutSceneXaxisCategoryorderMeanDescending LayoutSceneXaxisCategoryorder = "mean descending"
    LayoutSceneXaxisCategoryorderMedianAscending LayoutSceneXaxisCategoryorder = "median ascending"
    LayoutSceneXaxisCategoryorderMedianDescending LayoutSceneXaxisCategoryorder = "median descending"
    
)
// LayoutSceneXaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutSceneXaxisExponentformat string 

const (
    LayoutSceneXaxisExponentformatNone LayoutSceneXaxisExponentformat = "none"
    LayoutSceneXaxisExponentformatE1 LayoutSceneXaxisExponentformat = "e"
    LayoutSceneXaxisExponentformatE2 LayoutSceneXaxisExponentformat = "E"
    LayoutSceneXaxisExponentformatPower LayoutSceneXaxisExponentformat = "power"
    LayoutSceneXaxisExponentformatSi LayoutSceneXaxisExponentformat = "SI"
    LayoutSceneXaxisExponentformatB LayoutSceneXaxisExponentformat = "B"
    
)
// LayoutSceneXaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutSceneXaxisMirror interface{} 

var (
    LayoutSceneXaxisMirrorTrue LayoutSceneXaxisMirror = true
    LayoutSceneXaxisMirrorTicks LayoutSceneXaxisMirror = "ticks"
    LayoutSceneXaxisMirrorFalse LayoutSceneXaxisMirror = false
    LayoutSceneXaxisMirrorAll LayoutSceneXaxisMirror = "all"
    LayoutSceneXaxisMirrorAllticks LayoutSceneXaxisMirror = "allticks"
    
)
// LayoutSceneXaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutSceneXaxisRangemode string 

const (
    LayoutSceneXaxisRangemodeNormal LayoutSceneXaxisRangemode = "normal"
    LayoutSceneXaxisRangemodeTozero LayoutSceneXaxisRangemode = "tozero"
    LayoutSceneXaxisRangemodeNonnegative LayoutSceneXaxisRangemode = "nonnegative"
    
)
// LayoutSceneXaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutSceneXaxisShowexponent string 

const (
    LayoutSceneXaxisShowexponentAll LayoutSceneXaxisShowexponent = "all"
    LayoutSceneXaxisShowexponentFirst LayoutSceneXaxisShowexponent = "first"
    LayoutSceneXaxisShowexponentLast LayoutSceneXaxisShowexponent = "last"
    LayoutSceneXaxisShowexponentNone LayoutSceneXaxisShowexponent = "none"
    
)
// LayoutSceneXaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutSceneXaxisShowtickprefix string 

const (
    LayoutSceneXaxisShowtickprefixAll LayoutSceneXaxisShowtickprefix = "all"
    LayoutSceneXaxisShowtickprefixFirst LayoutSceneXaxisShowtickprefix = "first"
    LayoutSceneXaxisShowtickprefixLast LayoutSceneXaxisShowtickprefix = "last"
    LayoutSceneXaxisShowtickprefixNone LayoutSceneXaxisShowtickprefix = "none"
    
)
// LayoutSceneXaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutSceneXaxisShowticksuffix string 

const (
    LayoutSceneXaxisShowticksuffixAll LayoutSceneXaxisShowticksuffix = "all"
    LayoutSceneXaxisShowticksuffixFirst LayoutSceneXaxisShowticksuffix = "first"
    LayoutSceneXaxisShowticksuffixLast LayoutSceneXaxisShowticksuffix = "last"
    LayoutSceneXaxisShowticksuffixNone LayoutSceneXaxisShowticksuffix = "none"
    
)
// LayoutSceneXaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutSceneXaxisTickmode string 

const (
    LayoutSceneXaxisTickmodeAuto LayoutSceneXaxisTickmode = "auto"
    LayoutSceneXaxisTickmodeLinear LayoutSceneXaxisTickmode = "linear"
    LayoutSceneXaxisTickmodeArray LayoutSceneXaxisTickmode = "array"
    
)
// LayoutSceneXaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutSceneXaxisTicks string 

const (
    LayoutSceneXaxisTicksOutside LayoutSceneXaxisTicks = "outside"
    LayoutSceneXaxisTicksInside LayoutSceneXaxisTicks = "inside"
    LayoutSceneXaxisTicksEmpty LayoutSceneXaxisTicks = ""
    
)
// LayoutSceneXaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutSceneXaxisType string 

const (
    LayoutSceneXaxisTypeHyphenHyphen LayoutSceneXaxisType = "-"
    LayoutSceneXaxisTypeLinear LayoutSceneXaxisType = "linear"
    LayoutSceneXaxisTypeLog LayoutSceneXaxisType = "log"
    LayoutSceneXaxisTypeDate LayoutSceneXaxisType = "date"
    LayoutSceneXaxisTypeCategory LayoutSceneXaxisType = "category"
    
)
// LayoutSceneYaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutSceneYaxisAutorange interface{} 

var (
    LayoutSceneYaxisAutorangeTrue LayoutSceneYaxisAutorange = true
    LayoutSceneYaxisAutorangeFalse LayoutSceneYaxisAutorange = false
    LayoutSceneYaxisAutorangeReversed LayoutSceneYaxisAutorange = "reversed"
    
)
// LayoutSceneYaxisAutotypenumbers Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
type LayoutSceneYaxisAutotypenumbers string 

const (
    LayoutSceneYaxisAutotypenumbersConvertTypes LayoutSceneYaxisAutotypenumbers = "convert types"
    LayoutSceneYaxisAutotypenumbersStrict LayoutSceneYaxisAutotypenumbers = "strict"
    
)
// LayoutSceneYaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutSceneYaxisCalendar string 

const (
    LayoutSceneYaxisCalendarGregorian LayoutSceneYaxisCalendar = "gregorian"
    LayoutSceneYaxisCalendarChinese LayoutSceneYaxisCalendar = "chinese"
    LayoutSceneYaxisCalendarCoptic LayoutSceneYaxisCalendar = "coptic"
    LayoutSceneYaxisCalendarDiscworld LayoutSceneYaxisCalendar = "discworld"
    LayoutSceneYaxisCalendarEthiopian LayoutSceneYaxisCalendar = "ethiopian"
    LayoutSceneYaxisCalendarHebrew LayoutSceneYaxisCalendar = "hebrew"
    LayoutSceneYaxisCalendarIslamic LayoutSceneYaxisCalendar = "islamic"
    LayoutSceneYaxisCalendarJulian LayoutSceneYaxisCalendar = "julian"
    LayoutSceneYaxisCalendarMayan LayoutSceneYaxisCalendar = "mayan"
    LayoutSceneYaxisCalendarNanakshahi LayoutSceneYaxisCalendar = "nanakshahi"
    LayoutSceneYaxisCalendarNepali LayoutSceneYaxisCalendar = "nepali"
    LayoutSceneYaxisCalendarPersian LayoutSceneYaxisCalendar = "persian"
    LayoutSceneYaxisCalendarJalali LayoutSceneYaxisCalendar = "jalali"
    LayoutSceneYaxisCalendarTaiwan LayoutSceneYaxisCalendar = "taiwan"
    LayoutSceneYaxisCalendarThai LayoutSceneYaxisCalendar = "thai"
    LayoutSceneYaxisCalendarUmmalqura LayoutSceneYaxisCalendar = "ummalqura"
    
)
// LayoutSceneYaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutSceneYaxisCategoryorder string 

const (
    LayoutSceneYaxisCategoryorderTrace LayoutSceneYaxisCategoryorder = "trace"
    LayoutSceneYaxisCategoryorderCategoryAscending LayoutSceneYaxisCategoryorder = "category ascending"
    LayoutSceneYaxisCategoryorderCategoryDescending LayoutSceneYaxisCategoryorder = "category descending"
    LayoutSceneYaxisCategoryorderArray LayoutSceneYaxisCategoryorder = "array"
    LayoutSceneYaxisCategoryorderTotalAscending LayoutSceneYaxisCategoryorder = "total ascending"
    LayoutSceneYaxisCategoryorderTotalDescending LayoutSceneYaxisCategoryorder = "total descending"
    LayoutSceneYaxisCategoryorderMinAscending LayoutSceneYaxisCategoryorder = "min ascending"
    LayoutSceneYaxisCategoryorderMinDescending LayoutSceneYaxisCategoryorder = "min descending"
    LayoutSceneYaxisCategoryorderMaxAscending LayoutSceneYaxisCategoryorder = "max ascending"
    LayoutSceneYaxisCategoryorderMaxDescending LayoutSceneYaxisCategoryorder = "max descending"
    LayoutSceneYaxisCategoryorderSumAscending LayoutSceneYaxisCategoryorder = "sum ascending"
    LayoutSceneYaxisCategoryorderSumDescending LayoutSceneYaxisCategoryorder = "sum descending"
    LayoutSceneYaxisCategoryorderMeanAscending LayoutSceneYaxisCategoryorder = "mean ascending"
    LayoutSceneYaxisCategoryorderMeanDescending LayoutSceneYaxisCategoryorder = "mean descending"
    LayoutSceneYaxisCategoryorderMedianAscending LayoutSceneYaxisCategoryorder = "median ascending"
    LayoutSceneYaxisCategoryorderMedianDescending LayoutSceneYaxisCategoryorder = "median descending"
    
)
// LayoutSceneYaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutSceneYaxisExponentformat string 

const (
    LayoutSceneYaxisExponentformatNone LayoutSceneYaxisExponentformat = "none"
    LayoutSceneYaxisExponentformatE1 LayoutSceneYaxisExponentformat = "e"
    LayoutSceneYaxisExponentformatE2 LayoutSceneYaxisExponentformat = "E"
    LayoutSceneYaxisExponentformatPower LayoutSceneYaxisExponentformat = "power"
    LayoutSceneYaxisExponentformatSi LayoutSceneYaxisExponentformat = "SI"
    LayoutSceneYaxisExponentformatB LayoutSceneYaxisExponentformat = "B"
    
)
// LayoutSceneYaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutSceneYaxisMirror interface{} 

var (
    LayoutSceneYaxisMirrorTrue LayoutSceneYaxisMirror = true
    LayoutSceneYaxisMirrorTicks LayoutSceneYaxisMirror = "ticks"
    LayoutSceneYaxisMirrorFalse LayoutSceneYaxisMirror = false
    LayoutSceneYaxisMirrorAll LayoutSceneYaxisMirror = "all"
    LayoutSceneYaxisMirrorAllticks LayoutSceneYaxisMirror = "allticks"
    
)
// LayoutSceneYaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutSceneYaxisRangemode string 

const (
    LayoutSceneYaxisRangemodeNormal LayoutSceneYaxisRangemode = "normal"
    LayoutSceneYaxisRangemodeTozero LayoutSceneYaxisRangemode = "tozero"
    LayoutSceneYaxisRangemodeNonnegative LayoutSceneYaxisRangemode = "nonnegative"
    
)
// LayoutSceneYaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutSceneYaxisShowexponent string 

const (
    LayoutSceneYaxisShowexponentAll LayoutSceneYaxisShowexponent = "all"
    LayoutSceneYaxisShowexponentFirst LayoutSceneYaxisShowexponent = "first"
    LayoutSceneYaxisShowexponentLast LayoutSceneYaxisShowexponent = "last"
    LayoutSceneYaxisShowexponentNone LayoutSceneYaxisShowexponent = "none"
    
)
// LayoutSceneYaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutSceneYaxisShowtickprefix string 

const (
    LayoutSceneYaxisShowtickprefixAll LayoutSceneYaxisShowtickprefix = "all"
    LayoutSceneYaxisShowtickprefixFirst LayoutSceneYaxisShowtickprefix = "first"
    LayoutSceneYaxisShowtickprefixLast LayoutSceneYaxisShowtickprefix = "last"
    LayoutSceneYaxisShowtickprefixNone LayoutSceneYaxisShowtickprefix = "none"
    
)
// LayoutSceneYaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutSceneYaxisShowticksuffix string 

const (
    LayoutSceneYaxisShowticksuffixAll LayoutSceneYaxisShowticksuffix = "all"
    LayoutSceneYaxisShowticksuffixFirst LayoutSceneYaxisShowticksuffix = "first"
    LayoutSceneYaxisShowticksuffixLast LayoutSceneYaxisShowticksuffix = "last"
    LayoutSceneYaxisShowticksuffixNone LayoutSceneYaxisShowticksuffix = "none"
    
)
// LayoutSceneYaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutSceneYaxisTickmode string 

const (
    LayoutSceneYaxisTickmodeAuto LayoutSceneYaxisTickmode = "auto"
    LayoutSceneYaxisTickmodeLinear LayoutSceneYaxisTickmode = "linear"
    LayoutSceneYaxisTickmodeArray LayoutSceneYaxisTickmode = "array"
    
)
// LayoutSceneYaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutSceneYaxisTicks string 

const (
    LayoutSceneYaxisTicksOutside LayoutSceneYaxisTicks = "outside"
    LayoutSceneYaxisTicksInside LayoutSceneYaxisTicks = "inside"
    LayoutSceneYaxisTicksEmpty LayoutSceneYaxisTicks = ""
    
)
// LayoutSceneYaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutSceneYaxisType string 

const (
    LayoutSceneYaxisTypeHyphenHyphen LayoutSceneYaxisType = "-"
    LayoutSceneYaxisTypeLinear LayoutSceneYaxisType = "linear"
    LayoutSceneYaxisTypeLog LayoutSceneYaxisType = "log"
    LayoutSceneYaxisTypeDate LayoutSceneYaxisType = "date"
    LayoutSceneYaxisTypeCategory LayoutSceneYaxisType = "category"
    
)
// LayoutSceneZaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutSceneZaxisAutorange interface{} 

var (
    LayoutSceneZaxisAutorangeTrue LayoutSceneZaxisAutorange = true
    LayoutSceneZaxisAutorangeFalse LayoutSceneZaxisAutorange = false
    LayoutSceneZaxisAutorangeReversed LayoutSceneZaxisAutorange = "reversed"
    
)
// LayoutSceneZaxisAutotypenumbers Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
type LayoutSceneZaxisAutotypenumbers string 

const (
    LayoutSceneZaxisAutotypenumbersConvertTypes LayoutSceneZaxisAutotypenumbers = "convert types"
    LayoutSceneZaxisAutotypenumbersStrict LayoutSceneZaxisAutotypenumbers = "strict"
    
)
// LayoutSceneZaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutSceneZaxisCalendar string 

const (
    LayoutSceneZaxisCalendarGregorian LayoutSceneZaxisCalendar = "gregorian"
    LayoutSceneZaxisCalendarChinese LayoutSceneZaxisCalendar = "chinese"
    LayoutSceneZaxisCalendarCoptic LayoutSceneZaxisCalendar = "coptic"
    LayoutSceneZaxisCalendarDiscworld LayoutSceneZaxisCalendar = "discworld"
    LayoutSceneZaxisCalendarEthiopian LayoutSceneZaxisCalendar = "ethiopian"
    LayoutSceneZaxisCalendarHebrew LayoutSceneZaxisCalendar = "hebrew"
    LayoutSceneZaxisCalendarIslamic LayoutSceneZaxisCalendar = "islamic"
    LayoutSceneZaxisCalendarJulian LayoutSceneZaxisCalendar = "julian"
    LayoutSceneZaxisCalendarMayan LayoutSceneZaxisCalendar = "mayan"
    LayoutSceneZaxisCalendarNanakshahi LayoutSceneZaxisCalendar = "nanakshahi"
    LayoutSceneZaxisCalendarNepali LayoutSceneZaxisCalendar = "nepali"
    LayoutSceneZaxisCalendarPersian LayoutSceneZaxisCalendar = "persian"
    LayoutSceneZaxisCalendarJalali LayoutSceneZaxisCalendar = "jalali"
    LayoutSceneZaxisCalendarTaiwan LayoutSceneZaxisCalendar = "taiwan"
    LayoutSceneZaxisCalendarThai LayoutSceneZaxisCalendar = "thai"
    LayoutSceneZaxisCalendarUmmalqura LayoutSceneZaxisCalendar = "ummalqura"
    
)
// LayoutSceneZaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutSceneZaxisCategoryorder string 

const (
    LayoutSceneZaxisCategoryorderTrace LayoutSceneZaxisCategoryorder = "trace"
    LayoutSceneZaxisCategoryorderCategoryAscending LayoutSceneZaxisCategoryorder = "category ascending"
    LayoutSceneZaxisCategoryorderCategoryDescending LayoutSceneZaxisCategoryorder = "category descending"
    LayoutSceneZaxisCategoryorderArray LayoutSceneZaxisCategoryorder = "array"
    LayoutSceneZaxisCategoryorderTotalAscending LayoutSceneZaxisCategoryorder = "total ascending"
    LayoutSceneZaxisCategoryorderTotalDescending LayoutSceneZaxisCategoryorder = "total descending"
    LayoutSceneZaxisCategoryorderMinAscending LayoutSceneZaxisCategoryorder = "min ascending"
    LayoutSceneZaxisCategoryorderMinDescending LayoutSceneZaxisCategoryorder = "min descending"
    LayoutSceneZaxisCategoryorderMaxAscending LayoutSceneZaxisCategoryorder = "max ascending"
    LayoutSceneZaxisCategoryorderMaxDescending LayoutSceneZaxisCategoryorder = "max descending"
    LayoutSceneZaxisCategoryorderSumAscending LayoutSceneZaxisCategoryorder = "sum ascending"
    LayoutSceneZaxisCategoryorderSumDescending LayoutSceneZaxisCategoryorder = "sum descending"
    LayoutSceneZaxisCategoryorderMeanAscending LayoutSceneZaxisCategoryorder = "mean ascending"
    LayoutSceneZaxisCategoryorderMeanDescending LayoutSceneZaxisCategoryorder = "mean descending"
    LayoutSceneZaxisCategoryorderMedianAscending LayoutSceneZaxisCategoryorder = "median ascending"
    LayoutSceneZaxisCategoryorderMedianDescending LayoutSceneZaxisCategoryorder = "median descending"
    
)
// LayoutSceneZaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutSceneZaxisExponentformat string 

const (
    LayoutSceneZaxisExponentformatNone LayoutSceneZaxisExponentformat = "none"
    LayoutSceneZaxisExponentformatE1 LayoutSceneZaxisExponentformat = "e"
    LayoutSceneZaxisExponentformatE2 LayoutSceneZaxisExponentformat = "E"
    LayoutSceneZaxisExponentformatPower LayoutSceneZaxisExponentformat = "power"
    LayoutSceneZaxisExponentformatSi LayoutSceneZaxisExponentformat = "SI"
    LayoutSceneZaxisExponentformatB LayoutSceneZaxisExponentformat = "B"
    
)
// LayoutSceneZaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutSceneZaxisMirror interface{} 

var (
    LayoutSceneZaxisMirrorTrue LayoutSceneZaxisMirror = true
    LayoutSceneZaxisMirrorTicks LayoutSceneZaxisMirror = "ticks"
    LayoutSceneZaxisMirrorFalse LayoutSceneZaxisMirror = false
    LayoutSceneZaxisMirrorAll LayoutSceneZaxisMirror = "all"
    LayoutSceneZaxisMirrorAllticks LayoutSceneZaxisMirror = "allticks"
    
)
// LayoutSceneZaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutSceneZaxisRangemode string 

const (
    LayoutSceneZaxisRangemodeNormal LayoutSceneZaxisRangemode = "normal"
    LayoutSceneZaxisRangemodeTozero LayoutSceneZaxisRangemode = "tozero"
    LayoutSceneZaxisRangemodeNonnegative LayoutSceneZaxisRangemode = "nonnegative"
    
)
// LayoutSceneZaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutSceneZaxisShowexponent string 

const (
    LayoutSceneZaxisShowexponentAll LayoutSceneZaxisShowexponent = "all"
    LayoutSceneZaxisShowexponentFirst LayoutSceneZaxisShowexponent = "first"
    LayoutSceneZaxisShowexponentLast LayoutSceneZaxisShowexponent = "last"
    LayoutSceneZaxisShowexponentNone LayoutSceneZaxisShowexponent = "none"
    
)
// LayoutSceneZaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutSceneZaxisShowtickprefix string 

const (
    LayoutSceneZaxisShowtickprefixAll LayoutSceneZaxisShowtickprefix = "all"
    LayoutSceneZaxisShowtickprefixFirst LayoutSceneZaxisShowtickprefix = "first"
    LayoutSceneZaxisShowtickprefixLast LayoutSceneZaxisShowtickprefix = "last"
    LayoutSceneZaxisShowtickprefixNone LayoutSceneZaxisShowtickprefix = "none"
    
)
// LayoutSceneZaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutSceneZaxisShowticksuffix string 

const (
    LayoutSceneZaxisShowticksuffixAll LayoutSceneZaxisShowticksuffix = "all"
    LayoutSceneZaxisShowticksuffixFirst LayoutSceneZaxisShowticksuffix = "first"
    LayoutSceneZaxisShowticksuffixLast LayoutSceneZaxisShowticksuffix = "last"
    LayoutSceneZaxisShowticksuffixNone LayoutSceneZaxisShowticksuffix = "none"
    
)
// LayoutSceneZaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutSceneZaxisTickmode string 

const (
    LayoutSceneZaxisTickmodeAuto LayoutSceneZaxisTickmode = "auto"
    LayoutSceneZaxisTickmodeLinear LayoutSceneZaxisTickmode = "linear"
    LayoutSceneZaxisTickmodeArray LayoutSceneZaxisTickmode = "array"
    
)
// LayoutSceneZaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutSceneZaxisTicks string 

const (
    LayoutSceneZaxisTicksOutside LayoutSceneZaxisTicks = "outside"
    LayoutSceneZaxisTicksInside LayoutSceneZaxisTicks = "inside"
    LayoutSceneZaxisTicksEmpty LayoutSceneZaxisTicks = ""
    
)
// LayoutSceneZaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutSceneZaxisType string 

const (
    LayoutSceneZaxisTypeHyphenHyphen LayoutSceneZaxisType = "-"
    LayoutSceneZaxisTypeLinear LayoutSceneZaxisType = "linear"
    LayoutSceneZaxisTypeLog LayoutSceneZaxisType = "log"
    LayoutSceneZaxisTypeDate LayoutSceneZaxisType = "date"
    LayoutSceneZaxisTypeCategory LayoutSceneZaxisType = "category"
    
)
// LayoutSelectdirection When `dragmode` is set to *select*, this limits the selection of the drag to horizontal, vertical or diagonal. *h* only allows horizontal selection, *v* only vertical, *d* only diagonal and *any* sets no limit.
type LayoutSelectdirection string 

const (
    LayoutSelectdirectionH LayoutSelectdirection = "h"
    LayoutSelectdirectionV LayoutSelectdirection = "v"
    LayoutSelectdirectionD LayoutSelectdirection = "d"
    LayoutSelectdirectionAny LayoutSelectdirection = "any"
    
)
// LayoutTernaryAaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutTernaryAaxisExponentformat string 

const (
    LayoutTernaryAaxisExponentformatNone LayoutTernaryAaxisExponentformat = "none"
    LayoutTernaryAaxisExponentformatE1 LayoutTernaryAaxisExponentformat = "e"
    LayoutTernaryAaxisExponentformatE2 LayoutTernaryAaxisExponentformat = "E"
    LayoutTernaryAaxisExponentformatPower LayoutTernaryAaxisExponentformat = "power"
    LayoutTernaryAaxisExponentformatSi LayoutTernaryAaxisExponentformat = "SI"
    LayoutTernaryAaxisExponentformatB LayoutTernaryAaxisExponentformat = "B"
    
)
// LayoutTernaryAaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutTernaryAaxisLayer string 

const (
    LayoutTernaryAaxisLayerAboveTraces LayoutTernaryAaxisLayer = "above traces"
    LayoutTernaryAaxisLayerBelowTraces LayoutTernaryAaxisLayer = "below traces"
    
)
// LayoutTernaryAaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutTernaryAaxisShowexponent string 

const (
    LayoutTernaryAaxisShowexponentAll LayoutTernaryAaxisShowexponent = "all"
    LayoutTernaryAaxisShowexponentFirst LayoutTernaryAaxisShowexponent = "first"
    LayoutTernaryAaxisShowexponentLast LayoutTernaryAaxisShowexponent = "last"
    LayoutTernaryAaxisShowexponentNone LayoutTernaryAaxisShowexponent = "none"
    
)
// LayoutTernaryAaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutTernaryAaxisShowtickprefix string 

const (
    LayoutTernaryAaxisShowtickprefixAll LayoutTernaryAaxisShowtickprefix = "all"
    LayoutTernaryAaxisShowtickprefixFirst LayoutTernaryAaxisShowtickprefix = "first"
    LayoutTernaryAaxisShowtickprefixLast LayoutTernaryAaxisShowtickprefix = "last"
    LayoutTernaryAaxisShowtickprefixNone LayoutTernaryAaxisShowtickprefix = "none"
    
)
// LayoutTernaryAaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutTernaryAaxisShowticksuffix string 

const (
    LayoutTernaryAaxisShowticksuffixAll LayoutTernaryAaxisShowticksuffix = "all"
    LayoutTernaryAaxisShowticksuffixFirst LayoutTernaryAaxisShowticksuffix = "first"
    LayoutTernaryAaxisShowticksuffixLast LayoutTernaryAaxisShowticksuffix = "last"
    LayoutTernaryAaxisShowticksuffixNone LayoutTernaryAaxisShowticksuffix = "none"
    
)
// LayoutTernaryAaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutTernaryAaxisTickmode string 

const (
    LayoutTernaryAaxisTickmodeAuto LayoutTernaryAaxisTickmode = "auto"
    LayoutTernaryAaxisTickmodeLinear LayoutTernaryAaxisTickmode = "linear"
    LayoutTernaryAaxisTickmodeArray LayoutTernaryAaxisTickmode = "array"
    
)
// LayoutTernaryAaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutTernaryAaxisTicks string 

const (
    LayoutTernaryAaxisTicksOutside LayoutTernaryAaxisTicks = "outside"
    LayoutTernaryAaxisTicksInside LayoutTernaryAaxisTicks = "inside"
    LayoutTernaryAaxisTicksEmpty LayoutTernaryAaxisTicks = ""
    
)
// LayoutTernaryBaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutTernaryBaxisExponentformat string 

const (
    LayoutTernaryBaxisExponentformatNone LayoutTernaryBaxisExponentformat = "none"
    LayoutTernaryBaxisExponentformatE1 LayoutTernaryBaxisExponentformat = "e"
    LayoutTernaryBaxisExponentformatE2 LayoutTernaryBaxisExponentformat = "E"
    LayoutTernaryBaxisExponentformatPower LayoutTernaryBaxisExponentformat = "power"
    LayoutTernaryBaxisExponentformatSi LayoutTernaryBaxisExponentformat = "SI"
    LayoutTernaryBaxisExponentformatB LayoutTernaryBaxisExponentformat = "B"
    
)
// LayoutTernaryBaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutTernaryBaxisLayer string 

const (
    LayoutTernaryBaxisLayerAboveTraces LayoutTernaryBaxisLayer = "above traces"
    LayoutTernaryBaxisLayerBelowTraces LayoutTernaryBaxisLayer = "below traces"
    
)
// LayoutTernaryBaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutTernaryBaxisShowexponent string 

const (
    LayoutTernaryBaxisShowexponentAll LayoutTernaryBaxisShowexponent = "all"
    LayoutTernaryBaxisShowexponentFirst LayoutTernaryBaxisShowexponent = "first"
    LayoutTernaryBaxisShowexponentLast LayoutTernaryBaxisShowexponent = "last"
    LayoutTernaryBaxisShowexponentNone LayoutTernaryBaxisShowexponent = "none"
    
)
// LayoutTernaryBaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutTernaryBaxisShowtickprefix string 

const (
    LayoutTernaryBaxisShowtickprefixAll LayoutTernaryBaxisShowtickprefix = "all"
    LayoutTernaryBaxisShowtickprefixFirst LayoutTernaryBaxisShowtickprefix = "first"
    LayoutTernaryBaxisShowtickprefixLast LayoutTernaryBaxisShowtickprefix = "last"
    LayoutTernaryBaxisShowtickprefixNone LayoutTernaryBaxisShowtickprefix = "none"
    
)
// LayoutTernaryBaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutTernaryBaxisShowticksuffix string 

const (
    LayoutTernaryBaxisShowticksuffixAll LayoutTernaryBaxisShowticksuffix = "all"
    LayoutTernaryBaxisShowticksuffixFirst LayoutTernaryBaxisShowticksuffix = "first"
    LayoutTernaryBaxisShowticksuffixLast LayoutTernaryBaxisShowticksuffix = "last"
    LayoutTernaryBaxisShowticksuffixNone LayoutTernaryBaxisShowticksuffix = "none"
    
)
// LayoutTernaryBaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutTernaryBaxisTickmode string 

const (
    LayoutTernaryBaxisTickmodeAuto LayoutTernaryBaxisTickmode = "auto"
    LayoutTernaryBaxisTickmodeLinear LayoutTernaryBaxisTickmode = "linear"
    LayoutTernaryBaxisTickmodeArray LayoutTernaryBaxisTickmode = "array"
    
)
// LayoutTernaryBaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutTernaryBaxisTicks string 

const (
    LayoutTernaryBaxisTicksOutside LayoutTernaryBaxisTicks = "outside"
    LayoutTernaryBaxisTicksInside LayoutTernaryBaxisTicks = "inside"
    LayoutTernaryBaxisTicksEmpty LayoutTernaryBaxisTicks = ""
    
)
// LayoutTernaryCaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutTernaryCaxisExponentformat string 

const (
    LayoutTernaryCaxisExponentformatNone LayoutTernaryCaxisExponentformat = "none"
    LayoutTernaryCaxisExponentformatE1 LayoutTernaryCaxisExponentformat = "e"
    LayoutTernaryCaxisExponentformatE2 LayoutTernaryCaxisExponentformat = "E"
    LayoutTernaryCaxisExponentformatPower LayoutTernaryCaxisExponentformat = "power"
    LayoutTernaryCaxisExponentformatSi LayoutTernaryCaxisExponentformat = "SI"
    LayoutTernaryCaxisExponentformatB LayoutTernaryCaxisExponentformat = "B"
    
)
// LayoutTernaryCaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutTernaryCaxisLayer string 

const (
    LayoutTernaryCaxisLayerAboveTraces LayoutTernaryCaxisLayer = "above traces"
    LayoutTernaryCaxisLayerBelowTraces LayoutTernaryCaxisLayer = "below traces"
    
)
// LayoutTernaryCaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutTernaryCaxisShowexponent string 

const (
    LayoutTernaryCaxisShowexponentAll LayoutTernaryCaxisShowexponent = "all"
    LayoutTernaryCaxisShowexponentFirst LayoutTernaryCaxisShowexponent = "first"
    LayoutTernaryCaxisShowexponentLast LayoutTernaryCaxisShowexponent = "last"
    LayoutTernaryCaxisShowexponentNone LayoutTernaryCaxisShowexponent = "none"
    
)
// LayoutTernaryCaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutTernaryCaxisShowtickprefix string 

const (
    LayoutTernaryCaxisShowtickprefixAll LayoutTernaryCaxisShowtickprefix = "all"
    LayoutTernaryCaxisShowtickprefixFirst LayoutTernaryCaxisShowtickprefix = "first"
    LayoutTernaryCaxisShowtickprefixLast LayoutTernaryCaxisShowtickprefix = "last"
    LayoutTernaryCaxisShowtickprefixNone LayoutTernaryCaxisShowtickprefix = "none"
    
)
// LayoutTernaryCaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutTernaryCaxisShowticksuffix string 

const (
    LayoutTernaryCaxisShowticksuffixAll LayoutTernaryCaxisShowticksuffix = "all"
    LayoutTernaryCaxisShowticksuffixFirst LayoutTernaryCaxisShowticksuffix = "first"
    LayoutTernaryCaxisShowticksuffixLast LayoutTernaryCaxisShowticksuffix = "last"
    LayoutTernaryCaxisShowticksuffixNone LayoutTernaryCaxisShowticksuffix = "none"
    
)
// LayoutTernaryCaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutTernaryCaxisTickmode string 

const (
    LayoutTernaryCaxisTickmodeAuto LayoutTernaryCaxisTickmode = "auto"
    LayoutTernaryCaxisTickmodeLinear LayoutTernaryCaxisTickmode = "linear"
    LayoutTernaryCaxisTickmodeArray LayoutTernaryCaxisTickmode = "array"
    
)
// LayoutTernaryCaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutTernaryCaxisTicks string 

const (
    LayoutTernaryCaxisTicksOutside LayoutTernaryCaxisTicks = "outside"
    LayoutTernaryCaxisTicksInside LayoutTernaryCaxisTicks = "inside"
    LayoutTernaryCaxisTicksEmpty LayoutTernaryCaxisTicks = ""
    
)
// LayoutTitleXanchor Sets the title's horizontal alignment with respect to its x position. *left* means that the title starts at x, *right* means that the title ends at x and *center* means that the title's center is at x. *auto* divides `xref` by three and calculates the `xanchor` value automatically based on the value of `x`.
type LayoutTitleXanchor string 

const (
    LayoutTitleXanchorAuto LayoutTitleXanchor = "auto"
    LayoutTitleXanchorLeft LayoutTitleXanchor = "left"
    LayoutTitleXanchorCenter LayoutTitleXanchor = "center"
    LayoutTitleXanchorRight LayoutTitleXanchor = "right"
    
)
// LayoutTitleXref Sets the container `x` refers to. *container* spans the entire `width` of the plot. *paper* refers to the width of the plotting area only.
type LayoutTitleXref string 

const (
    LayoutTitleXrefContainer LayoutTitleXref = "container"
    LayoutTitleXrefPaper LayoutTitleXref = "paper"
    
)
// LayoutTitleYanchor Sets the title's vertical alignment with respect to its y position. *top* means that the title's cap line is at y, *bottom* means that the title's baseline is at y and *middle* means that the title's midline is at y. *auto* divides `yref` by three and calculates the `yanchor` value automatically based on the value of `y`.
type LayoutTitleYanchor string 

const (
    LayoutTitleYanchorAuto LayoutTitleYanchor = "auto"
    LayoutTitleYanchorTop LayoutTitleYanchor = "top"
    LayoutTitleYanchorMiddle LayoutTitleYanchor = "middle"
    LayoutTitleYanchorBottom LayoutTitleYanchor = "bottom"
    
)
// LayoutTitleYref Sets the container `y` refers to. *container* spans the entire `height` of the plot. *paper* refers to the height of the plotting area only.
type LayoutTitleYref string 

const (
    LayoutTitleYrefContainer LayoutTitleYref = "container"
    LayoutTitleYrefPaper LayoutTitleYref = "paper"
    
)
// LayoutTransitionEasing The easing function used for the transition
type LayoutTransitionEasing string 

const (
    LayoutTransitionEasingLinear LayoutTransitionEasing = "linear"
    LayoutTransitionEasingQuad LayoutTransitionEasing = "quad"
    LayoutTransitionEasingCubic LayoutTransitionEasing = "cubic"
    LayoutTransitionEasingSin LayoutTransitionEasing = "sin"
    LayoutTransitionEasingExp LayoutTransitionEasing = "exp"
    LayoutTransitionEasingCircle LayoutTransitionEasing = "circle"
    LayoutTransitionEasingElastic LayoutTransitionEasing = "elastic"
    LayoutTransitionEasingBack LayoutTransitionEasing = "back"
    LayoutTransitionEasingBounce LayoutTransitionEasing = "bounce"
    LayoutTransitionEasingLinearIn LayoutTransitionEasing = "linear-in"
    LayoutTransitionEasingQuadIn LayoutTransitionEasing = "quad-in"
    LayoutTransitionEasingCubicIn LayoutTransitionEasing = "cubic-in"
    LayoutTransitionEasingSinIn LayoutTransitionEasing = "sin-in"
    LayoutTransitionEasingExpIn LayoutTransitionEasing = "exp-in"
    LayoutTransitionEasingCircleIn LayoutTransitionEasing = "circle-in"
    LayoutTransitionEasingElasticIn LayoutTransitionEasing = "elastic-in"
    LayoutTransitionEasingBackIn LayoutTransitionEasing = "back-in"
    LayoutTransitionEasingBounceIn LayoutTransitionEasing = "bounce-in"
    LayoutTransitionEasingLinearOut LayoutTransitionEasing = "linear-out"
    LayoutTransitionEasingQuadOut LayoutTransitionEasing = "quad-out"
    LayoutTransitionEasingCubicOut LayoutTransitionEasing = "cubic-out"
    LayoutTransitionEasingSinOut LayoutTransitionEasing = "sin-out"
    LayoutTransitionEasingExpOut LayoutTransitionEasing = "exp-out"
    LayoutTransitionEasingCircleOut LayoutTransitionEasing = "circle-out"
    LayoutTransitionEasingElasticOut LayoutTransitionEasing = "elastic-out"
    LayoutTransitionEasingBackOut LayoutTransitionEasing = "back-out"
    LayoutTransitionEasingBounceOut LayoutTransitionEasing = "bounce-out"
    LayoutTransitionEasingLinearInOut LayoutTransitionEasing = "linear-in-out"
    LayoutTransitionEasingQuadInOut LayoutTransitionEasing = "quad-in-out"
    LayoutTransitionEasingCubicInOut LayoutTransitionEasing = "cubic-in-out"
    LayoutTransitionEasingSinInOut LayoutTransitionEasing = "sin-in-out"
    LayoutTransitionEasingExpInOut LayoutTransitionEasing = "exp-in-out"
    LayoutTransitionEasingCircleInOut LayoutTransitionEasing = "circle-in-out"
    LayoutTransitionEasingElasticInOut LayoutTransitionEasing = "elastic-in-out"
    LayoutTransitionEasingBackInOut LayoutTransitionEasing = "back-in-out"
    LayoutTransitionEasingBounceInOut LayoutTransitionEasing = "bounce-in-out"
    
)
// LayoutTransitionOrdering Determines whether the figure's layout or traces smoothly transitions during updates that make both traces and layout change.
type LayoutTransitionOrdering string 

const (
    LayoutTransitionOrderingLayoutFirst LayoutTransitionOrdering = "layout first"
    LayoutTransitionOrderingTracesFirst LayoutTransitionOrdering = "traces first"
    
)
// LayoutUniformtextMode Determines how the font size for various text elements are uniformed between each trace type. If the computed text sizes were smaller than the minimum size defined by `uniformtext.minsize` using *hide* option hides the text; and using *show* option shows the text without further downscaling. Please note that if the size defined by `minsize` is greater than the font size defined by trace, then the `minsize` is used.
type LayoutUniformtextMode interface{} 

var (
    LayoutUniformtextModeFalse LayoutUniformtextMode = false
    LayoutUniformtextModeHide LayoutUniformtextMode = "hide"
    LayoutUniformtextModeShow LayoutUniformtextMode = "show"
    
)
// LayoutXaxisAnchor If set to an opposite-letter axis id (e.g. `x2`, `y`), this axis is bound to the corresponding opposite-letter axis. If set to *free*, this axis' position is determined by `position`.
type LayoutXaxisAnchor string 

const (
    LayoutXaxisAnchorFree LayoutXaxisAnchor = "free"
    LayoutXaxisAnchorSlashCapexLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutXaxisAnchor = "/^x([2-9]|[1-9][0-9]+)?( domain)?$/"
    LayoutXaxisAnchorSlashCapeyLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutXaxisAnchor = "/^y([2-9]|[1-9][0-9]+)?( domain)?$/"
    
)
// LayoutXaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutXaxisAutorange interface{} 

var (
    LayoutXaxisAutorangeTrue LayoutXaxisAutorange = true
    LayoutXaxisAutorangeFalse LayoutXaxisAutorange = false
    LayoutXaxisAutorangeReversed LayoutXaxisAutorange = "reversed"
    
)
// LayoutXaxisAutotypenumbers Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
type LayoutXaxisAutotypenumbers string 

const (
    LayoutXaxisAutotypenumbersConvertTypes LayoutXaxisAutotypenumbers = "convert types"
    LayoutXaxisAutotypenumbersStrict LayoutXaxisAutotypenumbers = "strict"
    
)
// LayoutXaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutXaxisCalendar string 

const (
    LayoutXaxisCalendarGregorian LayoutXaxisCalendar = "gregorian"
    LayoutXaxisCalendarChinese LayoutXaxisCalendar = "chinese"
    LayoutXaxisCalendarCoptic LayoutXaxisCalendar = "coptic"
    LayoutXaxisCalendarDiscworld LayoutXaxisCalendar = "discworld"
    LayoutXaxisCalendarEthiopian LayoutXaxisCalendar = "ethiopian"
    LayoutXaxisCalendarHebrew LayoutXaxisCalendar = "hebrew"
    LayoutXaxisCalendarIslamic LayoutXaxisCalendar = "islamic"
    LayoutXaxisCalendarJulian LayoutXaxisCalendar = "julian"
    LayoutXaxisCalendarMayan LayoutXaxisCalendar = "mayan"
    LayoutXaxisCalendarNanakshahi LayoutXaxisCalendar = "nanakshahi"
    LayoutXaxisCalendarNepali LayoutXaxisCalendar = "nepali"
    LayoutXaxisCalendarPersian LayoutXaxisCalendar = "persian"
    LayoutXaxisCalendarJalali LayoutXaxisCalendar = "jalali"
    LayoutXaxisCalendarTaiwan LayoutXaxisCalendar = "taiwan"
    LayoutXaxisCalendarThai LayoutXaxisCalendar = "thai"
    LayoutXaxisCalendarUmmalqura LayoutXaxisCalendar = "ummalqura"
    
)
// LayoutXaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutXaxisCategoryorder string 

const (
    LayoutXaxisCategoryorderTrace LayoutXaxisCategoryorder = "trace"
    LayoutXaxisCategoryorderCategoryAscending LayoutXaxisCategoryorder = "category ascending"
    LayoutXaxisCategoryorderCategoryDescending LayoutXaxisCategoryorder = "category descending"
    LayoutXaxisCategoryorderArray LayoutXaxisCategoryorder = "array"
    LayoutXaxisCategoryorderTotalAscending LayoutXaxisCategoryorder = "total ascending"
    LayoutXaxisCategoryorderTotalDescending LayoutXaxisCategoryorder = "total descending"
    LayoutXaxisCategoryorderMinAscending LayoutXaxisCategoryorder = "min ascending"
    LayoutXaxisCategoryorderMinDescending LayoutXaxisCategoryorder = "min descending"
    LayoutXaxisCategoryorderMaxAscending LayoutXaxisCategoryorder = "max ascending"
    LayoutXaxisCategoryorderMaxDescending LayoutXaxisCategoryorder = "max descending"
    LayoutXaxisCategoryorderSumAscending LayoutXaxisCategoryorder = "sum ascending"
    LayoutXaxisCategoryorderSumDescending LayoutXaxisCategoryorder = "sum descending"
    LayoutXaxisCategoryorderMeanAscending LayoutXaxisCategoryorder = "mean ascending"
    LayoutXaxisCategoryorderMeanDescending LayoutXaxisCategoryorder = "mean descending"
    LayoutXaxisCategoryorderMedianAscending LayoutXaxisCategoryorder = "median ascending"
    LayoutXaxisCategoryorderMedianDescending LayoutXaxisCategoryorder = "median descending"
    
)
// LayoutXaxisConstrain If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines how that happens: by increasing the *range*, or by decreasing the *domain*. Default is *domain* for axes containing image traces, *range* otherwise.
type LayoutXaxisConstrain string 

const (
    LayoutXaxisConstrainRange LayoutXaxisConstrain = "range"
    LayoutXaxisConstrainDomain LayoutXaxisConstrain = "domain"
    
)
// LayoutXaxisConstraintoward If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines which direction we push the originally specified plot area. Options are *left*, *center* (default), and *right* for x axes, and *top*, *middle* (default), and *bottom* for y axes.
type LayoutXaxisConstraintoward string 

const (
    LayoutXaxisConstraintowardLeft LayoutXaxisConstraintoward = "left"
    LayoutXaxisConstraintowardCenter LayoutXaxisConstraintoward = "center"
    LayoutXaxisConstraintowardRight LayoutXaxisConstraintoward = "right"
    LayoutXaxisConstraintowardTop LayoutXaxisConstraintoward = "top"
    LayoutXaxisConstraintowardMiddle LayoutXaxisConstraintoward = "middle"
    LayoutXaxisConstraintowardBottom LayoutXaxisConstraintoward = "bottom"
    
)
// LayoutXaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutXaxisExponentformat string 

const (
    LayoutXaxisExponentformatNone LayoutXaxisExponentformat = "none"
    LayoutXaxisExponentformatE1 LayoutXaxisExponentformat = "e"
    LayoutXaxisExponentformatE2 LayoutXaxisExponentformat = "E"
    LayoutXaxisExponentformatPower LayoutXaxisExponentformat = "power"
    LayoutXaxisExponentformatSi LayoutXaxisExponentformat = "SI"
    LayoutXaxisExponentformatB LayoutXaxisExponentformat = "B"
    
)
// LayoutXaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutXaxisLayer string 

const (
    LayoutXaxisLayerAboveTraces LayoutXaxisLayer = "above traces"
    LayoutXaxisLayerBelowTraces LayoutXaxisLayer = "below traces"
    
)
// LayoutXaxisMatches If set to another axis id (e.g. `x2`, `y`), the range of this axis will match the range of the corresponding axis in data-coordinates space. Moreover, matching axes share auto-range values, category lists and histogram auto-bins. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden. Moreover, note that matching axes must have the same `type`.
type LayoutXaxisMatches string 

const (
    LayoutXaxisMatchesSlashCapexLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutXaxisMatches = "/^x([2-9]|[1-9][0-9]+)?( domain)?$/"
    LayoutXaxisMatchesSlashCapeyLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutXaxisMatches = "/^y([2-9]|[1-9][0-9]+)?( domain)?$/"
    
)
// LayoutXaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutXaxisMirror interface{} 

var (
    LayoutXaxisMirrorTrue LayoutXaxisMirror = true
    LayoutXaxisMirrorTicks LayoutXaxisMirror = "ticks"
    LayoutXaxisMirrorFalse LayoutXaxisMirror = false
    LayoutXaxisMirrorAll LayoutXaxisMirror = "all"
    LayoutXaxisMirrorAllticks LayoutXaxisMirror = "allticks"
    
)
// LayoutXaxisOverlaying If set a same-letter axis id, this axis is overlaid on top of the corresponding same-letter axis, with traces and axes visible for both axes. If *false*, this axis does not overlay any same-letter axes. In this case, for axes with overlapping domains only the highest-numbered axis will be visible.
type LayoutXaxisOverlaying string 

const (
    LayoutXaxisOverlayingFree LayoutXaxisOverlaying = "free"
    LayoutXaxisOverlayingSlashCapexLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutXaxisOverlaying = "/^x([2-9]|[1-9][0-9]+)?( domain)?$/"
    LayoutXaxisOverlayingSlashCapeyLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutXaxisOverlaying = "/^y([2-9]|[1-9][0-9]+)?( domain)?$/"
    
)
// LayoutXaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutXaxisRangemode string 

const (
    LayoutXaxisRangemodeNormal LayoutXaxisRangemode = "normal"
    LayoutXaxisRangemodeTozero LayoutXaxisRangemode = "tozero"
    LayoutXaxisRangemodeNonnegative LayoutXaxisRangemode = "nonnegative"
    
)
// LayoutXaxisRangeselectorXanchor Sets the range selector's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the range selector.
type LayoutXaxisRangeselectorXanchor string 

const (
    LayoutXaxisRangeselectorXanchorAuto LayoutXaxisRangeselectorXanchor = "auto"
    LayoutXaxisRangeselectorXanchorLeft LayoutXaxisRangeselectorXanchor = "left"
    LayoutXaxisRangeselectorXanchorCenter LayoutXaxisRangeselectorXanchor = "center"
    LayoutXaxisRangeselectorXanchorRight LayoutXaxisRangeselectorXanchor = "right"
    
)
// LayoutXaxisRangeselectorYanchor Sets the range selector's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the range selector.
type LayoutXaxisRangeselectorYanchor string 

const (
    LayoutXaxisRangeselectorYanchorAuto LayoutXaxisRangeselectorYanchor = "auto"
    LayoutXaxisRangeselectorYanchorTop LayoutXaxisRangeselectorYanchor = "top"
    LayoutXaxisRangeselectorYanchorMiddle LayoutXaxisRangeselectorYanchor = "middle"
    LayoutXaxisRangeselectorYanchorBottom LayoutXaxisRangeselectorYanchor = "bottom"
    
)
// LayoutXaxisRangesliderYaxisRangemode Determines whether or not the range of this axis in the rangeslider use the same value than in the main plot when zooming in/out. If *auto*, the autorange will be used. If *fixed*, the `range` is used. If *match*, the current range of the corresponding y-axis on the main subplot is used.
type LayoutXaxisRangesliderYaxisRangemode string 

const (
    LayoutXaxisRangesliderYaxisRangemodeAuto LayoutXaxisRangesliderYaxisRangemode = "auto"
    LayoutXaxisRangesliderYaxisRangemodeFixed LayoutXaxisRangesliderYaxisRangemode = "fixed"
    LayoutXaxisRangesliderYaxisRangemodeMatch LayoutXaxisRangesliderYaxisRangemode = "match"
    
)
// LayoutXaxisScaleanchor If set to another axis id (e.g. `x2`, `y`), the range of this axis changes together with the range of the corresponding axis such that the scale of pixels per unit is in a constant ratio. Both axes are still zoomable, but when you zoom one, the other will zoom the same amount, keeping a fixed midpoint. `constrain` and `constraintoward` determine how we enforce the constraint. You can chain these, ie `yaxis: {scaleanchor: *x*}, xaxis2: {scaleanchor: *y*}` but you can only link axes of the same `type`. The linked axis can have the opposite letter (to constrain the aspect ratio) or the same letter (to match scales across subplots). Loops (`yaxis: {scaleanchor: *x*}, xaxis: {scaleanchor: *y*}` or longer) are redundant and the last constraint encountered will be ignored to avoid possible inconsistent constraints via `scaleratio`. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden.
type LayoutXaxisScaleanchor string 

const (
    LayoutXaxisScaleanchorSlashCapexLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutXaxisScaleanchor = "/^x([2-9]|[1-9][0-9]+)?( domain)?$/"
    LayoutXaxisScaleanchorSlashCapeyLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutXaxisScaleanchor = "/^y([2-9]|[1-9][0-9]+)?( domain)?$/"
    
)
// LayoutXaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutXaxisShowexponent string 

const (
    LayoutXaxisShowexponentAll LayoutXaxisShowexponent = "all"
    LayoutXaxisShowexponentFirst LayoutXaxisShowexponent = "first"
    LayoutXaxisShowexponentLast LayoutXaxisShowexponent = "last"
    LayoutXaxisShowexponentNone LayoutXaxisShowexponent = "none"
    
)
// LayoutXaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutXaxisShowtickprefix string 

const (
    LayoutXaxisShowtickprefixAll LayoutXaxisShowtickprefix = "all"
    LayoutXaxisShowtickprefixFirst LayoutXaxisShowtickprefix = "first"
    LayoutXaxisShowtickprefixLast LayoutXaxisShowtickprefix = "last"
    LayoutXaxisShowtickprefixNone LayoutXaxisShowtickprefix = "none"
    
)
// LayoutXaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutXaxisShowticksuffix string 

const (
    LayoutXaxisShowticksuffixAll LayoutXaxisShowticksuffix = "all"
    LayoutXaxisShowticksuffixFirst LayoutXaxisShowticksuffix = "first"
    LayoutXaxisShowticksuffixLast LayoutXaxisShowticksuffix = "last"
    LayoutXaxisShowticksuffixNone LayoutXaxisShowticksuffix = "none"
    
)
// LayoutXaxisSide Determines whether a x (y) axis is positioned at the *bottom* (*left*) or *top* (*right*) of the plotting area.
type LayoutXaxisSide string 

const (
    LayoutXaxisSideTop LayoutXaxisSide = "top"
    LayoutXaxisSideBottom LayoutXaxisSide = "bottom"
    LayoutXaxisSideLeft LayoutXaxisSide = "left"
    LayoutXaxisSideRight LayoutXaxisSide = "right"
    
)
// LayoutXaxisSpikesnap Determines whether spikelines are stuck to the cursor or to the closest datapoints.
type LayoutXaxisSpikesnap string 

const (
    LayoutXaxisSpikesnapData LayoutXaxisSpikesnap = "data"
    LayoutXaxisSpikesnapCursor LayoutXaxisSpikesnap = "cursor"
    LayoutXaxisSpikesnapHoveredData LayoutXaxisSpikesnap = "hovered data"
    
)
// LayoutXaxisTicklabelmode Determines where tick labels are drawn with respect to their corresponding ticks and grid lines. Only has an effect for axes of `type` *date* When set to *period*, tick labels are drawn in the middle of the period between ticks.
type LayoutXaxisTicklabelmode string 

const (
    LayoutXaxisTicklabelmodeInstant LayoutXaxisTicklabelmode = "instant"
    LayoutXaxisTicklabelmodePeriod LayoutXaxisTicklabelmode = "period"
    
)
// LayoutXaxisTicklabelposition Determines where tick labels are drawn with respect to the axis Please note that top or bottom has no effect on x axes or when `ticklabelmode` is set to *period*. Similarly left or right has no effect on y axes or when `ticklabelmode` is set to *period*. Has no effect on *multicategory* axes or when `tickson` is set to *boundaries*. When used on axes linked by `matches` or `scaleanchor`, no extra padding for inside labels would be added by autorange, so that the scales could match.
type LayoutXaxisTicklabelposition string 

const (
    LayoutXaxisTicklabelpositionOutside LayoutXaxisTicklabelposition = "outside"
    LayoutXaxisTicklabelpositionInside LayoutXaxisTicklabelposition = "inside"
    LayoutXaxisTicklabelpositionOutsideTop LayoutXaxisTicklabelposition = "outside top"
    LayoutXaxisTicklabelpositionInsideTop LayoutXaxisTicklabelposition = "inside top"
    LayoutXaxisTicklabelpositionOutsideLeft LayoutXaxisTicklabelposition = "outside left"
    LayoutXaxisTicklabelpositionInsideLeft LayoutXaxisTicklabelposition = "inside left"
    LayoutXaxisTicklabelpositionOutsideRight LayoutXaxisTicklabelposition = "outside right"
    LayoutXaxisTicklabelpositionInsideRight LayoutXaxisTicklabelposition = "inside right"
    LayoutXaxisTicklabelpositionOutsideBottom LayoutXaxisTicklabelposition = "outside bottom"
    LayoutXaxisTicklabelpositionInsideBottom LayoutXaxisTicklabelposition = "inside bottom"
    
)
// LayoutXaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutXaxisTickmode string 

const (
    LayoutXaxisTickmodeAuto LayoutXaxisTickmode = "auto"
    LayoutXaxisTickmodeLinear LayoutXaxisTickmode = "linear"
    LayoutXaxisTickmodeArray LayoutXaxisTickmode = "array"
    
)
// LayoutXaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutXaxisTicks string 

const (
    LayoutXaxisTicksOutside LayoutXaxisTicks = "outside"
    LayoutXaxisTicksInside LayoutXaxisTicks = "inside"
    LayoutXaxisTicksEmpty LayoutXaxisTicks = ""
    
)
// LayoutXaxisTickson Determines where ticks and grid lines are drawn with respect to their corresponding tick labels. Only has an effect for axes of `type` *category* or *multicategory*. When set to *boundaries*, ticks and grid lines are drawn half a category to the left/bottom of labels.
type LayoutXaxisTickson string 

const (
    LayoutXaxisTicksonLabels LayoutXaxisTickson = "labels"
    LayoutXaxisTicksonBoundaries LayoutXaxisTickson = "boundaries"
    
)
// LayoutXaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutXaxisType string 

const (
    LayoutXaxisTypeHyphenHyphen LayoutXaxisType = "-"
    LayoutXaxisTypeLinear LayoutXaxisType = "linear"
    LayoutXaxisTypeLog LayoutXaxisType = "log"
    LayoutXaxisTypeDate LayoutXaxisType = "date"
    LayoutXaxisTypeCategory LayoutXaxisType = "category"
    LayoutXaxisTypeMulticategory LayoutXaxisType = "multicategory"
    
)
// LayoutYaxisAnchor If set to an opposite-letter axis id (e.g. `x2`, `y`), this axis is bound to the corresponding opposite-letter axis. If set to *free*, this axis' position is determined by `position`.
type LayoutYaxisAnchor string 

const (
    LayoutYaxisAnchorFree LayoutYaxisAnchor = "free"
    LayoutYaxisAnchorSlashCapexLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutYaxisAnchor = "/^x([2-9]|[1-9][0-9]+)?( domain)?$/"
    LayoutYaxisAnchorSlashCapeyLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutYaxisAnchor = "/^y([2-9]|[1-9][0-9]+)?( domain)?$/"
    
)
// LayoutYaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutYaxisAutorange interface{} 

var (
    LayoutYaxisAutorangeTrue LayoutYaxisAutorange = true
    LayoutYaxisAutorangeFalse LayoutYaxisAutorange = false
    LayoutYaxisAutorangeReversed LayoutYaxisAutorange = "reversed"
    
)
// LayoutYaxisAutotypenumbers Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
type LayoutYaxisAutotypenumbers string 

const (
    LayoutYaxisAutotypenumbersConvertTypes LayoutYaxisAutotypenumbers = "convert types"
    LayoutYaxisAutotypenumbersStrict LayoutYaxisAutotypenumbers = "strict"
    
)
// LayoutYaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutYaxisCalendar string 

const (
    LayoutYaxisCalendarGregorian LayoutYaxisCalendar = "gregorian"
    LayoutYaxisCalendarChinese LayoutYaxisCalendar = "chinese"
    LayoutYaxisCalendarCoptic LayoutYaxisCalendar = "coptic"
    LayoutYaxisCalendarDiscworld LayoutYaxisCalendar = "discworld"
    LayoutYaxisCalendarEthiopian LayoutYaxisCalendar = "ethiopian"
    LayoutYaxisCalendarHebrew LayoutYaxisCalendar = "hebrew"
    LayoutYaxisCalendarIslamic LayoutYaxisCalendar = "islamic"
    LayoutYaxisCalendarJulian LayoutYaxisCalendar = "julian"
    LayoutYaxisCalendarMayan LayoutYaxisCalendar = "mayan"
    LayoutYaxisCalendarNanakshahi LayoutYaxisCalendar = "nanakshahi"
    LayoutYaxisCalendarNepali LayoutYaxisCalendar = "nepali"
    LayoutYaxisCalendarPersian LayoutYaxisCalendar = "persian"
    LayoutYaxisCalendarJalali LayoutYaxisCalendar = "jalali"
    LayoutYaxisCalendarTaiwan LayoutYaxisCalendar = "taiwan"
    LayoutYaxisCalendarThai LayoutYaxisCalendar = "thai"
    LayoutYaxisCalendarUmmalqura LayoutYaxisCalendar = "ummalqura"
    
)
// LayoutYaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutYaxisCategoryorder string 

const (
    LayoutYaxisCategoryorderTrace LayoutYaxisCategoryorder = "trace"
    LayoutYaxisCategoryorderCategoryAscending LayoutYaxisCategoryorder = "category ascending"
    LayoutYaxisCategoryorderCategoryDescending LayoutYaxisCategoryorder = "category descending"
    LayoutYaxisCategoryorderArray LayoutYaxisCategoryorder = "array"
    LayoutYaxisCategoryorderTotalAscending LayoutYaxisCategoryorder = "total ascending"
    LayoutYaxisCategoryorderTotalDescending LayoutYaxisCategoryorder = "total descending"
    LayoutYaxisCategoryorderMinAscending LayoutYaxisCategoryorder = "min ascending"
    LayoutYaxisCategoryorderMinDescending LayoutYaxisCategoryorder = "min descending"
    LayoutYaxisCategoryorderMaxAscending LayoutYaxisCategoryorder = "max ascending"
    LayoutYaxisCategoryorderMaxDescending LayoutYaxisCategoryorder = "max descending"
    LayoutYaxisCategoryorderSumAscending LayoutYaxisCategoryorder = "sum ascending"
    LayoutYaxisCategoryorderSumDescending LayoutYaxisCategoryorder = "sum descending"
    LayoutYaxisCategoryorderMeanAscending LayoutYaxisCategoryorder = "mean ascending"
    LayoutYaxisCategoryorderMeanDescending LayoutYaxisCategoryorder = "mean descending"
    LayoutYaxisCategoryorderMedianAscending LayoutYaxisCategoryorder = "median ascending"
    LayoutYaxisCategoryorderMedianDescending LayoutYaxisCategoryorder = "median descending"
    
)
// LayoutYaxisConstrain If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines how that happens: by increasing the *range*, or by decreasing the *domain*. Default is *domain* for axes containing image traces, *range* otherwise.
type LayoutYaxisConstrain string 

const (
    LayoutYaxisConstrainRange LayoutYaxisConstrain = "range"
    LayoutYaxisConstrainDomain LayoutYaxisConstrain = "domain"
    
)
// LayoutYaxisConstraintoward If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines which direction we push the originally specified plot area. Options are *left*, *center* (default), and *right* for x axes, and *top*, *middle* (default), and *bottom* for y axes.
type LayoutYaxisConstraintoward string 

const (
    LayoutYaxisConstraintowardLeft LayoutYaxisConstraintoward = "left"
    LayoutYaxisConstraintowardCenter LayoutYaxisConstraintoward = "center"
    LayoutYaxisConstraintowardRight LayoutYaxisConstraintoward = "right"
    LayoutYaxisConstraintowardTop LayoutYaxisConstraintoward = "top"
    LayoutYaxisConstraintowardMiddle LayoutYaxisConstraintoward = "middle"
    LayoutYaxisConstraintowardBottom LayoutYaxisConstraintoward = "bottom"
    
)
// LayoutYaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutYaxisExponentformat string 

const (
    LayoutYaxisExponentformatNone LayoutYaxisExponentformat = "none"
    LayoutYaxisExponentformatE1 LayoutYaxisExponentformat = "e"
    LayoutYaxisExponentformatE2 LayoutYaxisExponentformat = "E"
    LayoutYaxisExponentformatPower LayoutYaxisExponentformat = "power"
    LayoutYaxisExponentformatSi LayoutYaxisExponentformat = "SI"
    LayoutYaxisExponentformatB LayoutYaxisExponentformat = "B"
    
)
// LayoutYaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutYaxisLayer string 

const (
    LayoutYaxisLayerAboveTraces LayoutYaxisLayer = "above traces"
    LayoutYaxisLayerBelowTraces LayoutYaxisLayer = "below traces"
    
)
// LayoutYaxisMatches If set to another axis id (e.g. `x2`, `y`), the range of this axis will match the range of the corresponding axis in data-coordinates space. Moreover, matching axes share auto-range values, category lists and histogram auto-bins. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden. Moreover, note that matching axes must have the same `type`.
type LayoutYaxisMatches string 

const (
    LayoutYaxisMatchesSlashCapexLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutYaxisMatches = "/^x([2-9]|[1-9][0-9]+)?( domain)?$/"
    LayoutYaxisMatchesSlashCapeyLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutYaxisMatches = "/^y([2-9]|[1-9][0-9]+)?( domain)?$/"
    
)
// LayoutYaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutYaxisMirror interface{} 

var (
    LayoutYaxisMirrorTrue LayoutYaxisMirror = true
    LayoutYaxisMirrorTicks LayoutYaxisMirror = "ticks"
    LayoutYaxisMirrorFalse LayoutYaxisMirror = false
    LayoutYaxisMirrorAll LayoutYaxisMirror = "all"
    LayoutYaxisMirrorAllticks LayoutYaxisMirror = "allticks"
    
)
// LayoutYaxisOverlaying If set a same-letter axis id, this axis is overlaid on top of the corresponding same-letter axis, with traces and axes visible for both axes. If *false*, this axis does not overlay any same-letter axes. In this case, for axes with overlapping domains only the highest-numbered axis will be visible.
type LayoutYaxisOverlaying string 

const (
    LayoutYaxisOverlayingFree LayoutYaxisOverlaying = "free"
    LayoutYaxisOverlayingSlashCapexLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutYaxisOverlaying = "/^x([2-9]|[1-9][0-9]+)?( domain)?$/"
    LayoutYaxisOverlayingSlashCapeyLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutYaxisOverlaying = "/^y([2-9]|[1-9][0-9]+)?( domain)?$/"
    
)
// LayoutYaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutYaxisRangemode string 

const (
    LayoutYaxisRangemodeNormal LayoutYaxisRangemode = "normal"
    LayoutYaxisRangemodeTozero LayoutYaxisRangemode = "tozero"
    LayoutYaxisRangemodeNonnegative LayoutYaxisRangemode = "nonnegative"
    
)
// LayoutYaxisScaleanchor If set to another axis id (e.g. `x2`, `y`), the range of this axis changes together with the range of the corresponding axis such that the scale of pixels per unit is in a constant ratio. Both axes are still zoomable, but when you zoom one, the other will zoom the same amount, keeping a fixed midpoint. `constrain` and `constraintoward` determine how we enforce the constraint. You can chain these, ie `yaxis: {scaleanchor: *x*}, xaxis2: {scaleanchor: *y*}` but you can only link axes of the same `type`. The linked axis can have the opposite letter (to constrain the aspect ratio) or the same letter (to match scales across subplots). Loops (`yaxis: {scaleanchor: *x*}, xaxis: {scaleanchor: *y*}` or longer) are redundant and the last constraint encountered will be ignored to avoid possible inconsistent constraints via `scaleratio`. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden.
type LayoutYaxisScaleanchor string 

const (
    LayoutYaxisScaleanchorSlashCapexLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutYaxisScaleanchor = "/^x([2-9]|[1-9][0-9]+)?( domain)?$/"
    LayoutYaxisScaleanchorSlashCapeyLparLbracket29RbracketOrLbracket19RbracketLbracket09RbracketPlusRparQuestionLparDomainRparQuestionDollarSlash LayoutYaxisScaleanchor = "/^y([2-9]|[1-9][0-9]+)?( domain)?$/"
    
)
// LayoutYaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutYaxisShowexponent string 

const (
    LayoutYaxisShowexponentAll LayoutYaxisShowexponent = "all"
    LayoutYaxisShowexponentFirst LayoutYaxisShowexponent = "first"
    LayoutYaxisShowexponentLast LayoutYaxisShowexponent = "last"
    LayoutYaxisShowexponentNone LayoutYaxisShowexponent = "none"
    
)
// LayoutYaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutYaxisShowtickprefix string 

const (
    LayoutYaxisShowtickprefixAll LayoutYaxisShowtickprefix = "all"
    LayoutYaxisShowtickprefixFirst LayoutYaxisShowtickprefix = "first"
    LayoutYaxisShowtickprefixLast LayoutYaxisShowtickprefix = "last"
    LayoutYaxisShowtickprefixNone LayoutYaxisShowtickprefix = "none"
    
)
// LayoutYaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutYaxisShowticksuffix string 

const (
    LayoutYaxisShowticksuffixAll LayoutYaxisShowticksuffix = "all"
    LayoutYaxisShowticksuffixFirst LayoutYaxisShowticksuffix = "first"
    LayoutYaxisShowticksuffixLast LayoutYaxisShowticksuffix = "last"
    LayoutYaxisShowticksuffixNone LayoutYaxisShowticksuffix = "none"
    
)
// LayoutYaxisSide Determines whether a x (y) axis is positioned at the *bottom* (*left*) or *top* (*right*) of the plotting area.
type LayoutYaxisSide string 

const (
    LayoutYaxisSideTop LayoutYaxisSide = "top"
    LayoutYaxisSideBottom LayoutYaxisSide = "bottom"
    LayoutYaxisSideLeft LayoutYaxisSide = "left"
    LayoutYaxisSideRight LayoutYaxisSide = "right"
    
)
// LayoutYaxisSpikesnap Determines whether spikelines are stuck to the cursor or to the closest datapoints.
type LayoutYaxisSpikesnap string 

const (
    LayoutYaxisSpikesnapData LayoutYaxisSpikesnap = "data"
    LayoutYaxisSpikesnapCursor LayoutYaxisSpikesnap = "cursor"
    LayoutYaxisSpikesnapHoveredData LayoutYaxisSpikesnap = "hovered data"
    
)
// LayoutYaxisTicklabelmode Determines where tick labels are drawn with respect to their corresponding ticks and grid lines. Only has an effect for axes of `type` *date* When set to *period*, tick labels are drawn in the middle of the period between ticks.
type LayoutYaxisTicklabelmode string 

const (
    LayoutYaxisTicklabelmodeInstant LayoutYaxisTicklabelmode = "instant"
    LayoutYaxisTicklabelmodePeriod LayoutYaxisTicklabelmode = "period"
    
)
// LayoutYaxisTicklabelposition Determines where tick labels are drawn with respect to the axis Please note that top or bottom has no effect on x axes or when `ticklabelmode` is set to *period*. Similarly left or right has no effect on y axes or when `ticklabelmode` is set to *period*. Has no effect on *multicategory* axes or when `tickson` is set to *boundaries*. When used on axes linked by `matches` or `scaleanchor`, no extra padding for inside labels would be added by autorange, so that the scales could match.
type LayoutYaxisTicklabelposition string 

const (
    LayoutYaxisTicklabelpositionOutside LayoutYaxisTicklabelposition = "outside"
    LayoutYaxisTicklabelpositionInside LayoutYaxisTicklabelposition = "inside"
    LayoutYaxisTicklabelpositionOutsideTop LayoutYaxisTicklabelposition = "outside top"
    LayoutYaxisTicklabelpositionInsideTop LayoutYaxisTicklabelposition = "inside top"
    LayoutYaxisTicklabelpositionOutsideLeft LayoutYaxisTicklabelposition = "outside left"
    LayoutYaxisTicklabelpositionInsideLeft LayoutYaxisTicklabelposition = "inside left"
    LayoutYaxisTicklabelpositionOutsideRight LayoutYaxisTicklabelposition = "outside right"
    LayoutYaxisTicklabelpositionInsideRight LayoutYaxisTicklabelposition = "inside right"
    LayoutYaxisTicklabelpositionOutsideBottom LayoutYaxisTicklabelposition = "outside bottom"
    LayoutYaxisTicklabelpositionInsideBottom LayoutYaxisTicklabelposition = "inside bottom"
    
)
// LayoutYaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutYaxisTickmode string 

const (
    LayoutYaxisTickmodeAuto LayoutYaxisTickmode = "auto"
    LayoutYaxisTickmodeLinear LayoutYaxisTickmode = "linear"
    LayoutYaxisTickmodeArray LayoutYaxisTickmode = "array"
    
)
// LayoutYaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutYaxisTicks string 

const (
    LayoutYaxisTicksOutside LayoutYaxisTicks = "outside"
    LayoutYaxisTicksInside LayoutYaxisTicks = "inside"
    LayoutYaxisTicksEmpty LayoutYaxisTicks = ""
    
)
// LayoutYaxisTickson Determines where ticks and grid lines are drawn with respect to their corresponding tick labels. Only has an effect for axes of `type` *category* or *multicategory*. When set to *boundaries*, ticks and grid lines are drawn half a category to the left/bottom of labels.
type LayoutYaxisTickson string 

const (
    LayoutYaxisTicksonLabels LayoutYaxisTickson = "labels"
    LayoutYaxisTicksonBoundaries LayoutYaxisTickson = "boundaries"
    
)
// LayoutYaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutYaxisType string 

const (
    LayoutYaxisTypeHyphenHyphen LayoutYaxisType = "-"
    LayoutYaxisTypeLinear LayoutYaxisType = "linear"
    LayoutYaxisTypeLog LayoutYaxisType = "log"
    LayoutYaxisTypeDate LayoutYaxisType = "date"
    LayoutYaxisTypeCategory LayoutYaxisType = "category"
    LayoutYaxisTypeMulticategory LayoutYaxisType = "multicategory"
    
)
// LayoutBarmode Determines how bars at the same location coordinate are displayed on the graph. With *stack*, the bars are stacked on top of one another With *relative*, the bars are stacked on top of one another, with negative values below the axis, positive values above With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars.
type LayoutBarmode string 

const (
    HistogramBarmodeStack LayoutBarmode = "stack"
    HistogramBarmodeGroup LayoutBarmode = "group"
    HistogramBarmodeOverlay LayoutBarmode = "overlay"
    HistogramBarmodeRelative LayoutBarmode = "relative"
    BarpolarBarmodeStack LayoutBarmode = "stack"
    BarpolarBarmodeOverlay LayoutBarmode = "overlay"
    BarBarmodeStack LayoutBarmode = "stack"
    BarBarmodeGroup LayoutBarmode = "group"
    BarBarmodeOverlay LayoutBarmode = "overlay"
    BarBarmodeRelative LayoutBarmode = "relative"
    
)
// LayoutBarnorm Sets the normalization for bar traces on the graph. With *fraction*, the value of each bar is divided by the sum of all values at that location coordinate. *percent* is the same but multiplied by 100 to show percentages.
type LayoutBarnorm string 

const (
    HistogramBarnormEmpty LayoutBarnorm = ""
    HistogramBarnormFraction LayoutBarnorm = "fraction"
    HistogramBarnormPercent LayoutBarnorm = "percent"
    BarBarnormEmpty LayoutBarnorm = ""
    BarBarnormFraction LayoutBarnorm = "fraction"
    BarBarnormPercent LayoutBarnorm = "percent"
    
)
// LayoutBoxmode Determines how boxes at the same location coordinate are displayed on the graph. If *group*, the boxes are plotted next to one another centered around the shared location. If *overlay*, the boxes are plotted over one another, you might need to set *opacity* to see them multiple boxes. Has no effect on traces that have *width* set.
type LayoutBoxmode string 

const (
    CandlestickBoxmodeGroup LayoutBoxmode = "group"
    CandlestickBoxmodeOverlay LayoutBoxmode = "overlay"
    BoxBoxmodeGroup LayoutBoxmode = "group"
    BoxBoxmodeOverlay LayoutBoxmode = "overlay"
    
)
// LayoutWaterfallmode Determines how bars at the same location coordinate are displayed on the graph. With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars.
type LayoutWaterfallmode string 

const (
    WaterfallWaterfallmodeGroup LayoutWaterfallmode = "group"
    WaterfallWaterfallmodeOverlay LayoutWaterfallmode = "overlay"
    
)
// LayoutViolinmode Determines how violins at the same location coordinate are displayed on the graph. If *group*, the violins are plotted next to one another centered around the shared location. If *overlay*, the violins are plotted over one another, you might need to set *opacity* to see them multiple violins. Has no effect on traces that have *width* set.
type LayoutViolinmode string 

const (
    ViolinViolinmodeGroup LayoutViolinmode = "group"
    ViolinViolinmodeOverlay LayoutViolinmode = "overlay"
    
)
// LayoutFunnelmode Determines how bars at the same location coordinate are displayed on the graph. With *stack*, the bars are stacked on top of one another With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars.
type LayoutFunnelmode string 

const (
    FunnelFunnelmodeStack LayoutFunnelmode = "stack"
    FunnelFunnelmodeGroup LayoutFunnelmode = "group"
    FunnelFunnelmodeOverlay LayoutFunnelmode = "overlay"
    
)
// LayoutClickmode Determines the mode of single click interactions. *event* is the default value and emits the `plotly_click` event. In addition this mode emits the `plotly_selected` event in drag modes *lasso* and *select*, but with no event data attached (kept for compatibility reasons). The *select* flag enables selecting single data points via click. This mode also supports persistent selections, meaning that pressing Shift while clicking, adds to / subtracts from an existing selection. *select* with `hovermode`: *x* can be confusing, consider explicitly setting `hovermode`: *closest* when using this feature. Selection events are sent accordingly as long as *event* flag is set as well. When the *event* flag is missing, `plotly_click` and `plotly_selected` events are not fired.
type LayoutClickmode string 

const (
    // Flags
    LayoutClickmodeEvent LayoutClickmode = "event"
    LayoutClickmodeSelect LayoutClickmode = "select"
    
    // Extra
    LayoutClickmodeNone LayoutClickmode = "none"
    
)
// LayoutLegendTraceorder Determines the order at which the legend items are displayed. If *normal*, the items are displayed top-to-bottom in the same order as the input data. If *reversed*, the items are displayed in the opposite order as *normal*. If *grouped*, the items are displayed in groups (when a trace `legendgroup` is provided). if *grouped+reversed*, the items are displayed in the opposite order as *grouped*.
type LayoutLegendTraceorder string 

const (
    // Flags
    LayoutLegendTraceorderReversed LayoutLegendTraceorder = "reversed"
    LayoutLegendTraceorderGrouped LayoutLegendTraceorder = "grouped"
    
    // Extra
    LayoutLegendTraceorderNormal LayoutLegendTraceorder = "normal"
    
)
// LayoutXaxisSpikemode Determines the drawing mode for the spike line If *toaxis*, the line is drawn from the data point to the axis the  series is plotted on. If *across*, the line is drawn across the entire plot area, and supercedes *toaxis*. If *marker*, then a marker dot is drawn on the axis the series is plotted on
type LayoutXaxisSpikemode string 

const (
    // Flags
    LayoutXaxisSpikemodeToaxis LayoutXaxisSpikemode = "toaxis"
    LayoutXaxisSpikemodeAcross LayoutXaxisSpikemode = "across"
    LayoutXaxisSpikemodeMarker LayoutXaxisSpikemode = "marker"
    
    // Extra
    
)
// LayoutYaxisSpikemode Determines the drawing mode for the spike line If *toaxis*, the line is drawn from the data point to the axis the  series is plotted on. If *across*, the line is drawn across the entire plot area, and supercedes *toaxis*. If *marker*, then a marker dot is drawn on the axis the series is plotted on
type LayoutYaxisSpikemode string 

const (
    // Flags
    LayoutYaxisSpikemodeToaxis LayoutYaxisSpikemode = "toaxis"
    LayoutYaxisSpikemodeAcross LayoutYaxisSpikemode = "across"
    LayoutYaxisSpikemodeMarker LayoutYaxisSpikemode = "marker"
    
    // Extra
    
)
