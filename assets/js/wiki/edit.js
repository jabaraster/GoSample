var PageEditor = React.createClass({displayName: "PageEditor",
    getInitialState: function() {
        return { title: this.props.title, body: '' };
    },
    componentDidMount: function() {
        $.get('/wiki/' + this.state.title + "/edit/data", null, function(pData) {
            this.setState(pData);
        }.bind(this))
    },
    handleSubmit: function() {
        var url = '/wiki/' + this.props.title
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
            React.createElement("div", {className: "PageEditor"}, 
                React.createElement("h1", null, this.state.title), 
                React.createElement("form", {onSubmit: this.handleSubmit, method: "post", action: '/wiki/' + this.state.title}, 
                    React.createElement("div", {className: "form-group"}, 
                        React.createElement("textarea", {className: "form-control", placeholder: "Body", ref: "body", value: this.state.body})
                    ), 
                    React.createElement("button", {className: "btn btn-primary"}, 
                        React.createElement("i", {className: "glyphicon glyphicon-ok"})
                    )
                )
            )
        )
    }
});

(function() {
    var title = $('#TitleValue').val();
    React.render(
        React.createElement(PageEditor, {title: title}),
        document.getElementById('PageEditor')
    );
})();
