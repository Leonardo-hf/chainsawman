import Guide from '@/components/Guide';
import {PageContainer} from '@ant-design/pro-components';
import styles from './index.less';

import {trim} from "@/utils/format";
import {useModel} from "@@/exports";


const HomePage: React.FC = (props) => {
    const {name} = useModel('global');
    return (
        <PageContainer ghost>
            <div className={styles.container}>
                <Guide name={trim(name)}/>
            </div>
        </PageContainer>
    );
};


export default HomePage

