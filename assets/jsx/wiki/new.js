var PageEditor = React.createClass({
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
            <form className="PageEditor" onSubmit={this.handleSubmit}>
                <div className="form-group">
                    <input type="text" className="form-control" placeholder="Title" ref="title" />
                </div>
                <div className="form-group">
                    <textarea className="form-control" placeholder="Body" ref="body"></textarea>
                </div>
                <button className="btn btn-primary">
                    <i className="glyphicon glyphicon-ok" />
                </button>
            </form>
        )
    }
});

React.render(
    <PageEditor />,
    document.getElementById('PageEditor')
);
