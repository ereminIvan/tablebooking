import React from "react";

var Button = React.createClass({
    defaultProps : {
        glyph: null
    },
    propTypes : {
        glyph: React.PropTypes.any
    },
    render : function() {
        var glyph = this.props.glyph
            ? <span className={"glyphicon glyphicon-" + this.props.glyph} aria-hidden="true"/>
            : null;
        return <button {...this.props} {...this.attr} aria-label="Left Align">{glyph}{this.props.children}</button>
    }
});

export default Button;