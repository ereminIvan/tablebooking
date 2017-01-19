import React from "react";
import Message from "./Message";

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


export default MessageSuccess;