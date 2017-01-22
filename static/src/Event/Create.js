import React from "react";
import PanelWrapper from "../Basic/PanelWrapper";
import {Header3} from "../Basic/Header";
import {Container} from "../Basic/Container";
import {Dialog} from "../Basic/Dialog";

class Create extends React.Component {
    render(){
        var dialogFooter =
            <div>
                <a type="button" className="btn btn-success btn-edit" href="/event/edit/">Продолжить редактирование</a>
                <a type="button" className="btn btn-success btn-list" href="/event/list">Список мероприятий</a>
                <button type="button" className="btn btn-success btn-reset" data-dismiss="modal">Создать еще одно</button>
            </div>;

        return(
            <Container>
                <Header3>Создание мероприятия</Header3>
                <PanelWrapper header="Создание нового мероприятия">
                    <form onSubmit={this.handleSubmit}>
                        <div className="form-group">
                            <label htmlFor="event_title">Название мероприятия</label>
                            <input name="event_title" type="text" className="form-control" id="event_title"
                                   aria-describedby="event_title_help" placeholder="Введите название мероприятия"/>
                            <small id="event_title_help" className="form-text text-muted"/>
                        </div>
                        <div className="form-group">
                            <label htmlFor="event_date_start">Дата начала мероприятия</label>
                            <div className="input-group date datetimepicker-input">
                                <input name="event_date_start" type="text" className="form-control" id="event_date_start"
                                       aria-describedby="event_title_help"/>
                                <span className="input-group-addon"><span className="glyphicon glyphicon-calendar"/></span>
                            </div>
                            <small id="event_date_start_help" className="form-text text-muted"/>
                        </div>
                        <button type="submit" className="btn btn-primary">Продолжить</button>
                    </form>
                </PanelWrapper>
                <Dialog header="Сохранение мероприятия" footer={dialogFooter}>Событие было сохранено</Dialog>
            </Container>
        )
    }
}

export default Create;