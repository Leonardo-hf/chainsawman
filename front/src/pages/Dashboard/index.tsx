import React from 'react';
import {useModel} from "@@/exports";
import SETable from "@/components/SETable";
import {Col, Divider, Row, Statistic} from "antd";
import CountUp from 'react-countup';
import {PageContainer, ProCard} from "@ant-design/pro-components";
import HHITable from "@/components/HHITable";
import {GraphRef2Group} from "@/models/global";
import {sum} from "@antfu/utils";

const Dashboard: React.FC = () => {
    const {initialState} = useModel('@@initialState')
    //@ts-ignore
    const {hotse, hhi} = initialState
    const {graphs} = useModel('global')
    const formatter = (value: number) => <CountUp end={value} separator=","/>

    return <PageContainer>
        <div style={{display: 'flex', flexDirection: 'column'}}>
            <Row gutter={16}>
                <Col span={8}>
                    <ProCard>
                        <Statistic title="图谱数目" value={graphs.length}
                            // @ts-ignore
                                   formatter={formatter}/>
                    </ProCard>
                </Col>
                <Col span={8}>
                    <ProCard>
                        <Statistic title="累计节点数目" value={sum(graphs.map((g: GraphRef2Group) => g.numNode))}
                            // @ts-ignore
                                   formatter={formatter}/>
                    </ProCard>
                </Col>
                <Col span={8}>
                    <ProCard>
                        <Statistic title="累计边数目" value={sum(graphs.map(g => g.numEdge))}
                            // @ts-ignore
                                   formatter={formatter}/>
                    </ProCard>
                </Col>
            </Row>
            <Divider/>
            <Row gutter={16}>
                {
                    hotse.map((s: Graph.HotSETopic) =>

                        <Col key={s.topic + s.language} span={12}>
                            <SETable hot={s}/>
                        </Col>)
                }
                {
                    hhi.map((h: Graph.HHILanguage) =>
                        <Col key={h.language} span={12}>
                            <HHITable hhi={h}/>
                        </Col>)
                }
            </Row>
        </div>
    </PageContainer>
};

export default Dashboard;

