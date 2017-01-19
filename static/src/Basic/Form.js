import React from "react";

var Form = React.createClass({
    render: function() {
        return <form {...this.attr} {...this.props}>{this.props.children}</form>
    }
});

var Group = React.createClass({
    render: function() {
        return <div className="form-group" {...this.attr}>{this.props.children}</div>
    }
});

var InputText = React.createClass({
    defaultProps : {
        name: null,
        label: null,
        placeholder: null,
        hint: null
    },
    propTypes : {
        name: React.PropTypes.string,
        label: React.PropTypes.string,
        placeholder: React.PropTypes.string,
        hint: React.PropTypes.string
    },
    render: function() {
        var label = this.props.label ? <label htmlFor={this.props.name}>{this.props.label}</label> : null;
        var hint = this.props.hint ? <small id={this.props.name + '_help'} className="form-text text-muted">{this.props.hint}</small> : null;
        return <div>
            {label}
            <input name={this.props.name} type="text"
                   className="form-control"
                   id={this.props.name}
                   aria-describedby={this.props.name + '_help'}
                   placeholder={this.props.placeholder}
            />
            {hint}
        </div>;
    }
});

var SelectGroup = React.createClass({
    render: function(){
        return null;
    }
});

export {Form, Group, InputText, SelectGroup}