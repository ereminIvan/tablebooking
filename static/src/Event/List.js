import React from "react";
import PanelWrapper from "../Basic.js";
import $ from "jquery";

var EventRows = React.createClass({
    getInitialState: function() {
        return {
            data: {}
        }
    },
    componentDidMount: function() {
        this.serverRequest = $.get(this.props.source, function (result) {
            this.setState({
                data: JSON.parse(result)
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
                    <a type="button" className="btn btn-info btn-sm"
                       onClick={self.actionEditEvent.bind(null, k)}>Изменить</a>
                    <a type="button" className="btn btn-danger btn-sm"
                       onClick={self.actionDeleteEvent.bind(null, k)}>Удалить</a>
                </td>
            </tr>
        });
        return <tbody>{renderRow}</tbody>
    }
});

var EventsTable = React.createClass({
    
    render : function () {
        const title = "Список событий";
        const header = "Список событий";
        const content = <div>
             <div><a href="/event/create" className="btn btn-success">Создать</a></div>
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
        </div>;
        return <PanelWrapper panelBodyContent={content}
                             panelHeadContent={title}
                             headerCaption={header}/>
    }
});

export default EventsTable;