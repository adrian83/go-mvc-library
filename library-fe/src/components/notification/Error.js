import React, { Component } from 'react';

class Error extends Component {

    constructor(props) { 
        super(props);

        this.hideError = this.hideError.bind(this);
    }

    hideError(errId){
        //console.log("hide", errId);
        const self = this;
        return function(event) {
            //console.log("real hide", errId);
            self.props.hideError(errId);
            event.preventDefault();
        }
    }

    renderError(id, message, details=[]){

        var i = 0;
        var detailsList = details.map(msg => (<div key={i++}>{msg}</div>));

        return (
            <div className="alert alert-danger alert-dismissible fade show" role="alert" key={id}>
                <div><b>{message}</b></div>
                {detailsList}
                <button type="button" className="close" aria-label="Close" onClick={this.hideError(id)}>
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>);
    }

    render() {

        var self = this;
        if(this.props && this.props.errors) {
            
            this.props.errors.forEach(function(error){ 
                console.log("error:", error);

                if(error.status) {

                    if(error.status === 500) {
                        return self.renderError(error.id, error.body);

                    } else if(error.status === 400) {
                        var msgs = error.body.map(v => v.message);
                        return self.renderError(error.id, "Invalida data", msgs);
                    
                    }  else if(error.status === 401) {
                        return self.renderError(error.id, error.body);

                    }  else if(error.status === 404) {
                        return self.renderError(error.id, error.body.message);
                    }

                    return self.renderError(error.id, error.body);
                }
                return self.renderError(new Date().getTime(), "Unknown error");
            });
        }
        return(<div></div>);
    }
}

export default Error;