import React from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { ContentHeader,Content, Header, Page, pageTheme } from '@backstage/core';
import {
  Container,
}from '@material-ui/core';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import Button from '@material-ui/core/Button';
import SearchIcon from '@material-ui/icons/Search';
import SaveAltIcon from '@material-ui/icons/SaveAlt';



// css style 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
    systemstlyeSave: {
        background: 'linear-gradient(45deg, #50E443 100%, #E8CBC0 100%)',
        border: 0,
        borderRadius: 3,
        boxShadow: '0 3px 5px 2px rgba(29, 43, 100, .1)',
        color: 'black',
        height: 48,
        padding: '0 30px',
        marginRight: theme.spacing(2),
    },
    
    systemstlyeSearch: {
        background: 'linear-gradient(45deg, #43E4A1 100%, #E8CBC0 100%)',
        border: 0,
        borderRadius: 3,
        boxShadow: '0 3px 5px 2px rgba(29, 43, 100, .1)',
        color: 'black',
        height: 48,
        padding: '0 30px',
        marginRight: theme.spacing(2),
    },



    systemstlye: {
      background: 'linear-gradient(45deg, #de1e47 100%, #E8CBC0 100%)',
      border: 0,
      borderRadius: 3,
      boxShadow: '0 3px 5px 2px rgba(29, 43, 100, .1)',
      color: 'black',
      height: 48,
      padding: '0 30px',
      marginRight: theme.spacing(2),
    },

  }),
);


  export default function Selectadminrepair(){
  const classes = useStyles();

  return (
    <Page theme={pageTheme.home}>
    <Header title="ระบบแจ้งซ่อมคอมพิวเตอร์">
    <Button 
        style={{ marginLeft: 20 }} 
        href="/"
        variant="contained"  
        startIcon={<ExitToAppIcon/>}
        > 
        ออกจากระบบ 
        </Button>
    </Header>
    <Content>
         <ContentHeader title="เลือกระบบที่ต้องการ">
    </ContentHeader>
      <Container maxWidth="sm">
      <Button style={{ marginBottom: 5 }}
            variant="contained"  href="/AdminrepairUI" size="large" fullWidth   startIcon={<SaveAltIcon/>} className={classes.systemstlyeSave}> บันทึกการซ่อมแซมคอมพิวเตอร์ของพนักงาน </Button>
      <Button style={{ marginBottom: 50 }}
            variant="contained"  href="/AdminrepairSearch" size="large" fullWidth   startIcon={<SearchIcon/>} className={classes.systemstlyeSearch}> ค้นหาหมายเลขบันทึกซ่อมแซมคอมพิวเตอร์ของพนักงาน </Button>
      <Button
            variant="contained"  href="/Group14" size="large" fullWidth   startIcon={<HomeRoundedIcon/>} className={classes.systemstlye}> กลับหน้าหลัก </Button>
      </Container>
    </Content>
  </Page>
  );
};