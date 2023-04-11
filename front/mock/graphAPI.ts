const graphs = [
    {name: "python", desc: "dependencies for pypi", nodes: 100, edges: 1000},
    {name: "java", desc: "dependencies for maven", nodes: 200, edges: 2000},
];

export default {
    'GET /api/graph/getAll': (req: any, res: any) => {
        res.json({
            graphs: graphs,
        });
    }
};
