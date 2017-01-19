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

export default Message;