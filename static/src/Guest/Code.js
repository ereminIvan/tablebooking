import React from "react";
import PanelWrapper from "../Basic/PanelWrapper";
import {Container} from "../Basic/Container";
import {Header3} from "../Basic/Header";

var Code = React.createClass({
    handleChangeCode: function(e) {
        console.log(e)
    },
    handleSubmit : function(e) {
        e.preventDefault();
        //todo return Guest.HandleCode($(this));
        return false;
    },
    render : function() {
        return (
            <Container>
                <Header3>Бронирование мест</Header3>
                <PanelWrapper header="Введите ваш код для продолжения бронирования">
                    <form onSubmit={(e) => this.handleSubmit(e)}>
                        <div className="form-group">
                            <label htmlFor="guest_code">Гостевой код</label>
                            <input value="" onChange={this.handleChangeCode}
                                   name="guest_code" type="text" className="form-control" id="guest_code"
                                   aria-describedby="guest_code_help" placeholder="Введите код"/>
                            <small id="guest_code_help" className="form-text text-muted" />
                        </div>
                        <button type="submit" className="btn btn-primary">Продолжить</button>
                    </form>
                </PanelWrapper>
            </Container>
        );
    }
});

export default Code