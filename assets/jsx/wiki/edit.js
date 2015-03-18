var PageEditor = React.createClass({
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
            <div className="PageEditor">
                <h1>{this.state.title}</h1>
                <form onSubmit={this.handleSubmit} method="post" action={'/wiki/' + this.state.title}>
                    <div className="form-group">
                        <textarea className="form-control" placeholder="Body" ref="body" value={this.state.body}></textarea>
                    </div>
                    <button className="btn btn-primary">
                        <i className="glyphicon glyphicon-ok" />
                    </button>
                </form>
            </div>
        )
    }
});

(function() {
    var title = $('#TitleValue').val();
    React.render(
        <PageEditor title={title} />,
        document.getElementById('PageEditor')
    );
})();
