import React from 'react';
import {useModel} from "@@/exports";
import SETable from "@/components/SETable";
import {Col, Divider, Row, Statistic} from "antd";
import CountUp from 'react-countup';
import {PageContainer, ProCard} from "@ant-design/pro-components";
import HHITable from "@/components/HHITable";

const Dashboard: React.FC = () => {
    const {initialState} = useModel('@@initialState')
    //@ts-ignore
    const {hotse, hhi} = initialState
    // const hotse: Graph.HotSETopic[] = [{
    //     language: 'Python',
    //     topic: 'breadth',
    //     software: [
    //         {
    //             artifact: 'requests',
    //             version: 'latest',
    //             score: 0.9727,
    //             homePage: 'https://github.com/psf/requests'
    //         },
    //         {
    //             artifact: 'numpy',
    //             version: 'latest',
    //             score: 0.8505,
    //             homePage: 'https://github.com/numpy/numpy'
    //         },
    //         {
    //             artifact: 'click',
    //             version: 'latest',
    //             score: 0.8290,
    //             homePage: 'https://github.com/pallets/click/issues/'
    //         },
    //         {
    //             artifact: 'pandas',
    //             version: 'latest',
    //             score: 0.8231,
    //             homePage: 'https://github.com/psf/requests'
    //         },
    //         {
    //             artifact: 'six',
    //             version: 'latest',
    //             score: 0.8015,
    //             homePage: 'https://github.com/pandas-dev/pandas'
    //         }],
    //     updateTime: new Date().getTime()
    // },
    //     {
    //         language: 'Python',
    //         topic: 'depth',
    //         software: [
    //             {
    //                 artifact: 'django',
    //                 version: 'latest',
    //                 score: 0.7354,
    //                 homePage: 'https://github.com/django/django'
    //             },
    //             {
    //                 artifact: 'jinja2',
    //                 version: 'latest',
    //                 score: 0.7327,
    //                 homePage: 'https://github.com/pallets/jinja/issues/'
    //             },
    //             {
    //                 artifact: 'flask',
    //                 version: 'latest',
    //                 score: 0.7015,
    //                 homePage: 'https://github.com/pallets/flask/issues/'
    //             },
    //             {
    //                 artifact: 'zope2',
    //                 version: 'latest',
    //                 score: 0.6453,
    //                 homePage: 'https://pypi.org/project/zope2'
    //             }],
    //         updateTime: new Date().getTime()
    //     },
    // ]
    const formatter = (value: number) => <CountUp end={value} separator=","/>

    return <PageContainer>
        <div style={{display: 'flex', flexDirection: 'column'}}>
            <Row gutter={16}>
                <Col span={8}>
                    <ProCard>
                        <Statistic title="图谱数目" value={4}
                            // @ts-ignore
                                   formatter={formatter}/>
                    </ProCard>
                </Col>
                <Col span={8}>
                    <ProCard>
                        <Statistic title="累计节点数目" value={4234} precision={2}
                            // @ts-ignore
                                   formatter={formatter}/>
                    </ProCard>
                </Col>
                <Col span={8}>
                    <ProCard>
                        <Statistic title="累计边数目" value={22321} precision={2}
                            // @ts-ignore
                                   formatter={formatter}/>
                    </ProCard>
                </Col>
            </Row>
            <Divider/>
            <Row gutter={16}>
                {
                    hotse.map((s: Graph.HotSETopic) =>
                        <Col span={12}>
                            <SETable hot={s}/>
                        </Col>)}
                {
                    hhi.map((h: Graph.HHILanguage) =>
                        <Col span={12}>
                            <HHITable hhi={h}/>
                        </Col>)
                }
            </Row>
        </div>
    </PageContainer>
};

export default Dashboard;

