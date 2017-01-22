import React, {Component} from "react";
import {Link} from "react-router";
import {Container} from "../components/Basic/Container";

export default class App extends Component {
    render() {
        return (
            <Container>
                <div className="col-md-2">
                    <ul className="nav nav-pills nav-stacked">
                        <li><Link to='/admin/event/create'>Event Create</Link></li>
                        <li><Link to='/admin/event/list'>Event List</Link></li>
                        <li><Link to='/guest/code'>Guest Code</Link></li>
                        <li><Link to='/admin/guest/create'>Guest Create</Link></li>
                    </ul>
                </div>
                <div className="col-md-10">
                    {this.props.children}
                </div>
            </Container>
        )
    }
}