import React from "react";

// var Button = React.createClass({
//     defaultProps : {
//         className: 'btn-default'
//     },
//     propTypes : {
//         className:  React.PropTypes.string
//     },
//     render : function() {
//         var className = "btn " + this.props.className;
//         return <button className={className}/>
//     }
// });

var PanelWrapper = React.createClass({
    defaultProps : {
        headerCaption: 'Default Header Caption',
        panelHeadContent: null,
        panelBodyContent: null
    },
    propTypes : {
        headerCaption: React.PropTypes.string,
        panelHeadContent: React.PropTypes.any,
        panelBodyContent: React.PropTypes.any
    },
    render : function() {
        return <div>
            <h3>{this.props.headerCaption}</h3>
            <div className="row">
                <div className="col-md-12">
                    <div className="panel panel-info">
                        <PanelHead content={this.props.panelHeadContent}/>
                        <PanelBody content={this.props.panelBodyContent} />
                    </div>
                </div>
            </div>
        </div>
    }
});

var PanelBody = React.createClass({
    defaultProps : {
        content: null
    },
    propTypes : {
        content : React.PropTypes.any
    },
    render : function() {
        return <div className="panel-body">{this.props.content}</div>
    }
});

var PanelHead = React.createClass({
    defaultProps : {
        content: null
    },
    propTypes : {
        content : React.PropTypes.any
    },
    render : function() {
        return <div className="panel-heading">{this.props.content}</div>
    }
});

export default PanelWrapper;