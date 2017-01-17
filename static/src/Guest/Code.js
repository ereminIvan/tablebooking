import React, {Component} from "react";
import PanelWrapper from "../Basic";

class Code extends Component {
    render() {
        const headerCaption = "Бронирование мест";
        const titleCaption = "Введите ваш код для продолжения бронирования";
        var content =
            <form onsubmit="return Guest.HandleCode($(this));">
                <div className="form-group">
                    <label for="guest_code">Гостевой код</label>
                    <input name="guest_code" type="text" className="form-control" id="guest_code"
                           aria-describedby="guest_code_help" placeholder="Введите код"/>
                    <small id="guest_code_help" className="form-text text-muted" />
                </div>
                <button type="submit" className="btn btn-primary">Продолжить</button>
            </form>;
        return (
            <PanelWrapper panelContent={content} headerCaption={headerCaption} titleCaption={titleCaption} />
        );
    }
}

export default Code