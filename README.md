# gin-demo
��gin��ܴ��һ����Ŀ�ṹ��������ٿ�����Ŀ��

### �ص�

- ����gorm������mysql�洢�����

- ����go redis�����ڲ�������

- ����uber/zap, zap��һ����Ч����־���

  ��Ŀ����־�����˷ֽ�:

  - request��־  

    ��¼http�����request��response���

  - panic��־

    ��¼http���������panic������־��������ٶ�λ��������

  - app��־

    ��¼���������У����Ǽ�¼��ҵ����־������������õ���־����־���м�¼��traceId��������ٸ�������鿴��ǰ�������־����
    ��־��ʽ����,������traceId��file��keywords���Լ����Ǽ�¼����Ҫ��Ϣ:
    
  ```shell
      2022-11-05 00:40:23.576	DEBUG	dao/user.go:24	83e51872-35a3-455d-9c5f-bd64ee140d4c	sql	[0.527ms] [rows:0] SELECT * FROM `users` WHERE `users`.`uid` = 7777 ORDER BY `users`.`uid` LIMIT 1


  ```
- ����gopkg.in/yaml.v3���û��������ǵ�������

### Ŀ¼�ṹ

- config -����
- handler -ҵ�������&·������
- logs -��־Ŀ¼
- middleware -�м��Ŀ¼
- entity -��ģ��Ŀ¼
- logic -ҵ���߼�Ŀ¼
- dao -���ݿ����
- svc -������log��db��redis�Ȼ�����������ķ�װ
- types �����������������������
- server.go -��������ļ�

### ����

- ��־������зָ�
- ҵ����־���Ը���traceId�鿴���������������־
- ��Ӧ�����ع̶����ֶΣ�����code��message��data��traceid