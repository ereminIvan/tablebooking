import React from "react";

class Header3 extends React.Component {
    render() {
        return(
            <h3>{this.props.children}</h3>
        )
    }
}

export {Header3};