var PageEditor = React.createClass({displayName: "PageEditor",
    handleSubmit: function() {
        var url = '/wiki/' + React.findDOMNode(this.refs.title).value;
        var body = React.findDOMNode(this.refs.body).value;
        $.ajax({
            url: url,
            type: 'post',
            data: { body: body },
            success: function() {
                location.href = url;
            },
            fail: function() {
                console.log(arguments);
            }
        });
        return false;
    },
    render: function() {
        return (
            React.createElement("form", {className: "PageEditor", onSubmit: this.handleSubmit}, 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("input", {type: "text", className: "form-control", placeholder: "Title", ref: "title"})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("textarea", {className: "form-control", placeholder: "Body", ref: "body"})
                ), 
                React.createElement("button", {className: "btn btn-primary"}, 
                    React.createElement("i", {className: "glyphicon glyphicon-ok"})
                )
            )
        )
    }
});

React.render(
    React.createElement(PageEditor, null),
    document.getElementById('PageEditor')
);
