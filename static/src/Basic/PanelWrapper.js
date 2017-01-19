import React from "react";

var PanelWrapper = React.createClass({
    defaultProps : {
        headerCaption: 'Default Header Caption',
        panelHeadContent: null
    },
    propTypes : {
        headerCaption: React.PropTypes.string,
        panelHeadContent: React.PropTypes.any
    },
    render : function() {
        return <div>
            <h3>{this.props.headerCaption}</h3>
            <div className="row">
                <div className="col-md-12">
                    <div className="panel panel-info">
                        <PanelHead>{this.props.panelHeadContent}</PanelHead>
                        <PanelBody>{this.props.children}</PanelBody>
                    </div>
                </div>
            </div>
        </div>
    }
});

var PanelBody = React.createClass({
    render : function() {
        return <div className="panel-body">{this.props.children}</div>
    }
});

var PanelHead = React.createClass({
    render : function() {
        return <div className="panel-heading">{this.props.children}</div>
    }
});

export default PanelWrapper;