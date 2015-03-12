var NewMessageForm = React.createClass({
    onSubmit: function(e) {
        e.preventDefault();
        var author = React.findDOMNode(this.refs.author).value.trim();
        var message = React.findDOMNode(this.refs.message).value.trim();
        alert('TODO データ送信 -> ' + author + ': ' + message);
        return false;
    },
    render: function() {
        return (
            <form className="NewMessageForm">
                <input type="text" placeholder="投稿者名" ref="author" />
                <input type="text" placeholder="メッセージ" ref="message" />
                <button className="btn btn-primary btn-lg" title="投稿" onClick={this.onSubmit}>
                    <i className="glyphicon glyphicon-cloud-upload"></i>
                </button>
            </form>
        )
    }
});

var Message = React.createClass({
    render: function() {
        return (
            <div className="Message">
                <label>{this.props.data.author}</label>
                <p>{this.props.data.message}</p>
            </div>
        )
    }
});

var Messages = React.createClass({
    render: function() {
        var nodes = this.props.data.map(function(pMessage) {
            return (<Message data={pMessage} />)
        });
        return (
            <div className="Messages">{nodes}</div>
        )
    }
});

React.render(
    <NewMessageForm />,
    document.getElementById('NewMessageForm')
);
var data = [
    {message: 'メッセージ1', author: 'jabaraster@gmail.com'},
    {message: 'メッセージ2', author: 'ah@jabara.info'}
];
React.render(
    <Messages data={data} />,
    document.getElementById('Messages')
);