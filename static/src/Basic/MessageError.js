import React from "react";
import Message from "./Message";

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

export default MessageError;