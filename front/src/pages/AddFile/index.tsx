import {PageContainer} from "@ant-design/pro-components";
import {Button, Dropdown, Form, Upload, UploadFile, UploadProps} from "antd";
import {EllipsisOutlined, InboxOutlined, UploadOutlined} from "@ant-design/icons";
import ProCard from "@ant-design/pro-card";
import {ProForm} from "@ant-design/pro-form/lib";
import { Input } from 'antd';
import {useState} from "react";

import {upload} from "@/services/file/file";
import axios from "axios";
import FormItem from "antd/es/form/FormItem";
import {values} from "@antv/util";
import Icon from "antd/es/icon";

const AddPage: React.FC = () => {
    const [fileList, setFileList] = useState<UploadFile[]>([]);
    const uploadProps: UploadProps = {

        beforeUpload: (file) => {
            setFileList([...fileList, file]);
            console.log(fileList)
            return false;
        },

        maxCount:2

    };

    const normFile = (e: any) => {
        return e?.fileList;
    };

    const handleFinish = (value: any) => {
        const { name,file } = value;
        console.log(name)
        console.log(file)
        const formData = new FormData();
        file.forEach((item, index) => {
            formData.append(`file-${index}`, item.originFileObj);
            console.log(item.originFileObj)
        })

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