import React from "react";
import PanelWrapper from "../Basic/PanelWrapper";
import {Form, Group} from "../Basic/Form.js";
import {Button} from "../Basic/Button.js";
import {Row} from "../Basic/Container";
import {Header3} from "../Basic/Header";

class Create extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            guest_name: "",
            guest_last_name: "",
            guest_is_vip: false,
            event_title: "",
            events: []
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    componentDidMount () {
        var self = this;
        fetch('http://localhost:8080/event/list')
            .then(
                function(response) {
                    if (response.status !== 200) {
                        console.log('Looks like there was a problem. Status Code: ' + response.status);
                        return;
                    }
                    response.json().then(function(data) {
                        self.setState({events:data.data});
                    });
                }
            )
            .catch(function(err) {
                console.log('Fetch Error :-S', err);
            });
    }

    handleChange(e) {
        var data = {};
        if (e.target.type === "checkbox") {
            data[e.target.name] = e.target.checked;
        } else {
            data[e.target.name] = e.target.value;
        }
        this.setState(data);
        console.log(this.state);
    }

    handleSubmit(e) {
        e.preventDefault();

        var self = this;
        var data = {
            guest_name: self.state.guest_name,
            guest_last_name: self.state.guest_last_name,
            guest_is_vip: self.state.guest_is_vip,
            event_title: self.state.event_title
        };
        fetch('http://localhost:8080/guest/create', {
            method: "POST",
            body: JSON.stringify(data)
        })
            .then(
                function(response) {
                    if (response.status !== 200) {
                        console.log('Looks like there was a problem. Status Code: ' + response.status);
                        return;
                    }
                    response.json().then(function(data) {
                        console.log(data);
                    });
                }
            )
            .catch(function(err) {
                console.log('Fetch Error :-S', err);
            });
    }

    render() {
        return (
            <Row>
                <Header3>Добавление нового гостя</Header3>
                <PanelWrapper header="Добавление нового гостя">
                    <Form onSubmit={this.handleSubmit}>
                        <Group>
                            <label htmlFor="guest_name">Имя гостя</label>
                            <input name="guest_name" type="text" className="form-control" id="guest_name" onChange={this.handleChange}
                                   aria-describedby="guest_name_help" placeholder="Введите Имя гостя" value={this.state.guest_name}/>
                            <small id="guest_name_help" className="form-text text-muted">Пожалуйста введите правильное Имя гостя</small>
                        </Group>
                        <Group>
                            <label htmlFor="guest_last_name">Фамилия гостя</label>
                            <input name="guest_last_name" type="text" className="form-control" id="guest_last_name" onChange={this.handleChange}
                                   aria-describedby="guest_last_name_help" placeholder="Введите Фамилию гостя" value={this.state.guest_last_name}/>
                            <small id="guest_last_name_help" className="form-text text-muted"/>
                        </Group>
                        <Group>
                            <label htmlFor="event_title">Мероприятие</label>
                            <select name="event_title" className="form-control" id="event_title" aria-describedby="event_title_help"
                                    value={this.state.event_title} onChange={this.handleChange}>
                                {this.state.events.map(function(e) {
                                    return <option key={e.id} value={e.id}>{e.title}</option>;
                                })}
                            </select>
                            <small id="event_title_help" className="form-text text-muted">Выберете мероприятие в которое добавляете гостя</small>
                        </Group>
                        <Group>
                            <div className="checkbox">
                                <label htmlFor="guest_is_vip">
                                    <input type="checkbox" name="guest_is_vip" value="" id="guest_is_vip"
                                           aria-describedby="guest_is_vip_help"
                                           onChange={this.handleChange}
                                           checked={this.state.guest_is_vip}/>VIP гость</label>
                            </div>
                        </Group>
                        <Button type="submit" className="btn btn-primary">Сохранить</Button>
                    </Form>
                </PanelWrapper>
            </Row>
        )
    }
}

export default Create;