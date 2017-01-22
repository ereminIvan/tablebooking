import React from "react";
import PanelWrapper from "../Basic/PanelWrapper";
import {Header3} from "../Basic/Header";
import {Row} from "../Basic/Container";
import {Dialog} from "../Basic/Dialog";
import {Form, Group} from "../Basic/Form";
import {Button} from "../Basic/Button";

class Create extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            event_title: "Default event title",
            event_date_start: new Date()
        };

        this.handleCreate = this.handleCreate.bind(this);
        this.handleChange = this.handleChange.bind(this);
    }

    handleCreate(e) {
        e.preventDefault(e);
        console.log("Dummy handle create with %s", this.state);
    }

    handleChange(event) {
        var data = {};
        data[event.target.name] = event.target.value;
        this.setState(data);
    }

    render(){
        var dialogFooter =
            <div>
                <Button className="btn btn-success btn-edit">Продолжить редактирование</Button>
                <Button className="btn btn-success btn-success btn-list">Список мероприятий</Button>
                <Button className="btn btn-success btn-reset" data-dismiss="modal">Создать еще одно</Button>
            </div>;

        return(
            <Row>
                <Header3>Создание мероприятия</Header3>
                <PanelWrapper header="Создание нового мероприятия">
                    <Form onSubmit={this.handleCreate}>
                        <Group>
                            <label htmlFor="event_title">Название мероприятия</label>
                            <input name="event_title" type="text" className="form-control" id="event_title"
                                   aria-describedby="event_title_help" placeholder="Введите название мероприятия"
                                   value={this.state.event_title} onChange={this.handleChange}
                            />
                            <small id="event_title_help" className="form-text text-muted"/>
                        </Group>
                        <Group>
                            <label htmlFor="event_date_start">Дата начала мероприятия</label>
                            <div className="input-group date datetimepicker-input">
                                <input name="event_date_start" type="text" className="form-control" id="event_date_start"
                                       aria-describedby="event_title_help"
                                       value={this.state.event_date_start.toLocaleDateString()} onChange={this.handleChange}
                                />
                                <span className="input-group-addon"><span className="glyphicon glyphicon-calendar"/></span>
                            </div>
                            <small id="event_date_start_help" className="form-text text-muted"/>
                        </Group>
                        <Button type="submit" className="btn btn-primary">Продолжить</Button>
                    </Form>
                </PanelWrapper>
                <Dialog header="Сохранение мероприятия" footer={dialogFooter}>Событие было сохранено</Dialog>
            </Row>
        )
    }
}

export default Create;