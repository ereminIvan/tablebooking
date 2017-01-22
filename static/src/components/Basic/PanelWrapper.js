import React from "react";

var PanelWrapper = React.createClass({
    defaultProps : {
        header: null
    },
    propTypes : {
        header: React.PropTypes.any
    },
    render : function() {
        return <div className="row">
                <div className="col-md-12">
                    <div className="panel panel-info">
                        <PanelHead>{this.props.header}</PanelHead>
                        <PanelBody>{this.props.children}</PanelBody>
                    </div>
                </div>
            </div>;
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