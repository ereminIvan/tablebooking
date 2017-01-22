import React from "react";

class Dialog extends React.Component {
    render(){
        return (
            <div className="modal fade" tabIndex="-1" role="dialog" id="dialog">
                <div className="modal-dialog" role="document">
                    <div className="modal-content">
                        <div className="modal-header">
                            <button type="button" className="close" data-dismiss="modal" aria-label="Close">
                                <span aria-hidden="true">&times;</span>
                            </button>
                            <h4 className="modal-title">{this.props.header}</h4>
                        </div>
                        <div className="modal-body">
                            <p>{this.props.children}</p>
                        </div>
                        <div className="modal-footer">{this.props.footer}</div>
                    </div>
                </div>
            </div>
        )
    }
}

export {Dialog};