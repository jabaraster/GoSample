var NewMessageForm = React.createClass({displayName: "NewMessageForm",
    onSubmit: function(e) {
        e.preventDefault();
        var author = React.findDOMNode(this.refs.author).value.trim();
        var message = React.findDOMNode(this.refs.message).value.trim();
        alert('TODO データ送信 -> ' + author + ': ' + message);
        return false;
    },
    render: function() {
        return (
            React.createElement("form", {className: "NewMessageForm"}, 
                React.createElement("input", {type: "text", placeholder: "投稿者名", ref: "author"}), 
                React.createElement("input", {type: "text", placeholder: "メッセージ", ref: "message"}), 
                React.createElement("button", {className: "btn btn-primary btn-lg", title: "投稿", onClick: this.onSubmit}, 
                    React.createElement("i", {className: "glyphicon glyphicon-cloud-upload"})
                )
            )
        )
    }
});

var Message = React.createClass({displayName: "Message",
    render: function() {
        return (
            React.createElement("div", {className: "Message"}, 
                React.createElement("label", null, this.props.data.author), 
                React.createElement("p", null, this.props.data.message)
            )
        )
    }
});

var Messages = React.createClass({displayName: "Messages",
    render: function() {
        var nodes = this.props.data.map(function(pMessage) {
            return (React.createElement(Message, {data: pMessage}))
        });
        return (
            React.createElement("div", {className: "Messages"}, nodes)
        )
    }
});

React.render(
    React.createElement(NewMessageForm, null),
    document.getElementById('NewMessageForm')
);
var data = [
    {message: 'メッセージ1', author: 'jabaraster@gmail.com'},
    {message: 'メッセージ2', author: 'ah@jabara.info'}
];
React.render(
    React.createElement(Messages, {data: data}),
    document.getElementById('Messages')
);