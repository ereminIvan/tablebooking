import React from "react";
import PanelWrapper from "../Basic/PanelWrapper";
import {Form, Group} from "../Basic/Form.js";
import {Button} from "../Basic/Button.js";
import {Container} from "../Basic/Container";
import {Header3} from "../Basic/Header";


class Create extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            guest_name: "",
            guest_last_name: "",
            guest_is_vip: "",
            event_title: ""
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    };
    handleChange(event) {
        var data = {};
        data[event.target.name] = event.target.value;
        console.log(this.state);
        this.setState(data);
    };
    handleSubmit(event) {
        event.preventDefault();
        console.log(this.state);
    };
    render(){
        return (
            <Container>
                <Header3>Добавление нового гостя на мероприятие</Header3>
                <PanelWrapper header="Добавление нового гостя">
                    <Form method="POST" action="/guest/create" onSubmit={this.handleSubmit}>
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
                                    value={this.state.value} onChange={this.handleChange}>
                                <option value="todo1">todo1</option>
                                <option value="todo2">todo2</option>
                            </select>
                            <small id="event_title_help" className="form-text text-muted">Выберете мероприятие на которое в которое добавляете гостя</small>
                        </Group>
                        <Group>
                            <div className="checkbox">
                                <label htmlFor="guest_is_vip"><input type="checkbox" name="guest_is_vip" value="" id="guest_is_vip"
                                                                     aria-describedby="guest_is_vip_help"/>VIP гость</label>
                            </div>
                        </Group>
                        <Button type="submit" className="btn btn-primary">Сохранить</Button>
                    </Form>
                </PanelWrapper>
            </Container>
        )
    }
}

export default Create;