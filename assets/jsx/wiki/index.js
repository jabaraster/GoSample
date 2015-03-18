var Page = React.createClass({
    handleClick: function(event) {
        location.href = '/wiki/' + this.props.page.title;
    },
    render: function() {
        return (
            <div className="Page" onClick={this.handleClick}>{this.props.page.title}</div>
        )
    }
});

var PageList = React.createClass({
    getInitialState: function() {
        return {data: []};
    },
    componentDidMount: function() {
        $.get('/wiki/index/data', null, function(pData) {
            this.setState({ data: pData });
        }.bind(this));
        var cont = $('.PageList');
        cont.on('mouseover', 'div.Page', function() {
            $(this).addClass('selected');
        });
        cont.on('mouseout', 'div.Page', function() {
            $(this).removeClass('selected');
        });
    },
    render: function() {
        var nodes = this.state.data.map(function(pPage) {
            return ( <Page page={pPage} /> )
        });
        return (
            <div className="PageList">
                <h1>ページ一覧</h1>
                <div>
                    <a className="btn btn-default" href="/wiki/new">
                        <i className="glyphicon glyphicon-plus" />
                    </a>
                </div>
                {nodes}
            </div>
        )
    }
});

React.render(
    <PageList />,
    document.getElementById('PageList')
);
