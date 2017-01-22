import React from "react";

var Message = React.createClass({
    defaultProps : {
        messageType: 'default'
    },
    propTypes : {
        messageType: React.PropTypes.string
    },
    render : function() {
        var messageType = "alert alert-" + this.props.messageType + " fade in alert-dismissable";
        return <div className={messageType} style="margin-top:18px;">
            <a href="#" className="close" data-dismiss="alert" aria-label="close" title="close">×</a>
            {this.props.children}
        </div>
    }
});

var MessageError = React.createClass({
    defaultProps : {
        messageType: 'error'
    },
    propTypes : {
        messageType: React.PropTypes.string
    },
    render : function() {
        return <Message messageType={this.props.messageType}>{this.props.children}</Message>
    }
});

var MessageSuccess = React.createClass({
    defaultProps : {
        messageType: 'success'
    },
    propTypes : {
        messageType: React.PropTypes.string
    },
    render : function() {
        return <Message messageType={this.props.messageType}>{this.props.children}</Message>
    }
});

export default {MessageError, MessageSuccess};