import React from "react";
import PanelWrapper from "../Basic/PanelWrapper";
import {Row} from "../Basic/Container";
import {Header3} from "../Basic/Header";
import {Form, Group} from "../Basic/Form";

class Code extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            code: ""
        };
        this.handleCheck = this.handleCheck.bind(this);
        this.handleChange = this.handleChange.bind(this);
    }

    handleChange(e) {
        this.setState({code: e.target.value});
    }

    handleCheck(e) {
        e.preventDefault();
        //todo implement code check
    }

    render(){
        return(
            <Row>
                <Header3>Бронирование мест</Header3>
                <PanelWrapper header="Введение гостевого кода">
                    <Form onSubmit={this.handleCheck}>
                        <Group>
                            <label htmlFor="guest_code">Гостевой код</label>
                            <input name="guest_code" type="text" className="form-control" id="guest_code"
                                   aria-describedby="guest_code_help" placeholder="Введите ваш код для продолжения бронирования"
                                   value={this.state.code} onChange={this.handleChange}
                            />
                            <small id="guest_code_help" className="form-text text-muted" />
                        </Group>
                        <button type="submit" className="btn btn-primary">Продолжить</button>
                    </Form>
                </PanelWrapper>
            </Row>
        );
    }
}

export default Code