import React from "react";
import PanelWrapper from "../Basic/PanelWrapper";
import ButtonA from "../Basic/ButtonA";
import Button from "../Basic/Button";
import $ from "jquery";

var EventRows = React.createClass({
    getInitialState: function() {
        return {
            data: {}
        }
    },
    componentDidMount: function() {
        this.serverRequest = $.get(this.props.source, function (result) {
            var d = JSON.parse(result);
            console.log(d);
            this.setState({
                data: d.data
            });
        }.bind(this));
    },
    componentWillUnmount: function() {
        this.serverRequest.abort();
    },
    actionEditEvent : function(d) {
        return false;
    },
    actionDeleteEvent : function(id) {
        $.post("/event/delete/" + id, function (result) {
            var data = JSON.parse(result);
            console.log(data);
        });
    },
    render: function() {
        var self = this;
        var renderRow = Object.keys(this.state.data).map(function (k, idx) {
            return <tr key={idx}>
                <td>{k}</td>
                <td>{self.state.data[k].title}</td>
                <td>{self.state.data[k].start_date}</td>
                <td>{self.state.data[k].tables}</td>
                <td>
                    <Button className="btn btn-info btn-sm" glyph="edit" onClick={self.actionEditEvent.bind(null, k)} />&nbsp;
                    <Button className="btn btn-danger btn-sm" glyph="remove" onClick={self.actionDeleteEvent.bind(null, k)} />
                </td>
            </tr>
        });
        return <tbody>{renderRow}</tbody>
    }
});

var EventsTable = React.createClass({
    
    render : function () {
        return <PanelWrapper panelHeadContent="Список событий" headerCaption="Список событий">
            <ButtonA className="btn btn-success" href="/event/create">Создать</ButtonA>
            <table className="table">
                <thead>
                <tr>
                    <th>#</th>
                    <th>Название мероприятия</th>
                    <th>Дата начала</th>
                    <th>Столы</th>
                    <th>Действия</th>
                </tr>
                </thead>
                <EventRows source="http://localhost:8080/event/list" />
            </table>
        </PanelWrapper>;
    }
});

export default EventsTable;