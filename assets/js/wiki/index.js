var Page = React.createClass({displayName: "Page",
    handleClick: function(event) {
        location.href = '/wiki/' + this.props.page.title;
    },
    render: function() {
        return (
            React.createElement("div", {className: "Page", onClick: this.handleClick}, this.props.page.title)
        )
    }
});

var PageList = React.createClass({displayName: "PageList",
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
            return ( React.createElement(Page, {page: pPage}) )
        });
        return (
            React.createElement("div", {className: "PageList"}, 
                React.createElement("h1", null, "ページ一覧"), 
                React.createElement("div", null, 
                    React.createElement("a", {className: "btn btn-default", href: "/wiki/new"}, 
                        React.createElement("i", {className: "glyphicon glyphicon-plus"})
                    )
                ), 
                nodes
            )
        )
    }
});

React.render(
    React.createElement(PageList, null),
    document.getElementById('PageList')
);
