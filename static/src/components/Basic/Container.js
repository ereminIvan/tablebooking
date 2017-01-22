import React from "react";

class Container extends React.Component {
    render(){
        return(
            <div className="container">{this.props.children}</div>
        )
    }
}

class ContainerFluid extends React.Component {
    render(){
        return (
            <div className="container-fluid">{this.props.children}</div>
        )
    }
}

class Row extends React.Component {
    render(){
        return (
            <div className="row">{this.props.children}</div>
        )
    }
}

export {Container, ContainerFluid, Row};