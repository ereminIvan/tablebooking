import React from "react";

var ButtonA  = React.createClass({
    render : function() {
        return <a {...this.props} {...this.attr}>{this.props.children}</a>
    }
});

export default ButtonA;