import {PageContainer} from "@ant-design/pro-components";
import {Button, Dropdown, Form, Upload, UploadFile, UploadProps} from "antd";
import {EllipsisOutlined, InboxOutlined, UploadOutlined} from "@ant-design/icons";
import ProCard from "@ant-design/pro-card";
import {ProForm} from "@ant-design/pro-form/lib";
import { Input } from 'antd';
import {useRef, useState} from "react";

import {upload} from "@/services/file/file";
import {createGraph} from "@/services/graph/graph";
import axios from "axios";
import FormItem from "antd/es/form/FormItem";
import {values} from "@antv/util";
import Icon from "antd/es/icon";

const AddPage: React.FC = () => {
    const [fileList, setFileList] = useState<UploadFile[]>([]);
    const [ifSuccess,setIfSuccess]=useState(false)
    const uploadProps: UploadProps = {

        beforeUpload: (file) => {
            return false;
        },

        maxCount:2

    };

    const normFile = (e: any) => {
        return e?.fileList;
    };

    const handleFinish = async (value: any) => {
        const {name, file} = value;
        const nodeData = new FormData();
        nodeData.append('node', file[0].originFileObj)
        //console.log(nodeData.get('node'))
        const edgeData = new FormData();
        edgeData.append('edge', file[1].originFileObj)
        let nodeId='', edgeId='';
        await upload(nodeData).then(res => {
            nodeId = res.id
        }).catch(e => console.log(e))
        await upload(edgeData).then(res => {
            edgeId = res.id
        }).catch(e => console.log(e))
        createGraph({taskId: 0, nodeId: nodeId,edgeId:edgeId,graph:name})
            .then(res=>{setIfSuccess(true)})
    }

    const uploadFile=async (options:any)=>{
        console.log(upload())
    }
    return (
        <PageContainer>
            <ProCard style={{height:"fit-content"}}>
                <Form onFinish={handleFinish} >
                    <FormItem name='name'>
                        <Input></Input>
                    </FormItem>
                    <Form.Item
                        name="file"
                        valuePropName="file"
                        getValueFromEvent={normFile}
                        label="文件"
                    >
                        <Upload.Dragger {...uploadProps}>
                            <p>
                                <InboxOutlined />
                            </p>
                            <p>点击或拖拽文件到上传区域</p>
                        </Upload.Dragger>
                    </Form.Item>

                    <Form.Item wrapperCol={{ span: 12, offset: 6 }}>
                        <Button type="primary" htmlType="submit">
                            确认发布
                        </Button>
                    </Form.Item>
                </Form>
            </ProCard>

        </PageContainer>
    )
}

export default AddPage