CREATE TABLE bira (
    idregrut INTEGER NOT NULL,
    idpik    INTEGER NOT NULL
);

CREATE UNIQUE INDEX bira__idx ON
    bira (
        idpik
    ASC );

CREATE UNIQUE INDEX bira__idxv1 ON
    bira (
        idregrut
    ASC );

ALTER TABLE bira ADD CONSTRAINT bira_pk PRIMARY KEY ( idregrut,
                                                      idpik );

CREATE TABLE draft (
    iddraft     INTEGER NOT NULL,
    gododrdraft VARCHAR2(4 CHAR) NOT NULL,
    lokodrdraft VARCHAR2(32 CHAR) NOT NULL
);

ALTER TABLE draft ADD CONSTRAINT draft_pk PRIMARY KEY ( iddraft );

CREATE TABLE igrac (
    id     INTEGER NOT NULL,
    visigr VARCHAR2(6 CHAR) NOT NULL,
    tezigr VARCHAR2(6 CHAR) NOT NULL,
    pozigr VARCHAR2(2) NOT NULL
);

ALTER TABLE igrac
    ADD CHECK ( pozigr IN ( 'C', 'PF', 'PG', 'SF', 'SG' ) );

ALTER TABLE igrac ADD CONSTRAINT igrac_pk PRIMARY KEY ( id );

CREATE TABLE imovinazatrgovinutima (
    idimotrgtim     INTEGER NOT NULL,
    datdodimotrgtim DATE NOT NULL,
    belesimotrgtim  VARCHAR2(64 CHAR) NOT NULL,
    idtipimotrg     INTEGER NOT NULL,
    idtim           INTEGER
);

CREATE UNIQUE INDEX imovinazatrgovinutima__idx ON
    imovinazatrgovinutima (
        idtim
    ASC );

ALTER TABLE imovinazatrgovinutima ADD CONSTRAINT imovinazatrgovinutima_pk PRIMARY KEY ( idimotrgtim );

CREATE TABLE intervju (
    idint     INTEGER NOT NULL,
    mesodrint VARCHAR2(32 CHAR) NOT NULL,
    datvreint DATE NOT NULL,
    belesint  VARCHAR2(64 CHAR) NOT NULL,
    idpozint  INTEGER NOT NULL,
    idregrut  INTEGER
);

CREATE UNIQUE INDEX intervju__idx ON
    intervju (
        idpozint
    ASC );

ALTER TABLE intervju ADD CONSTRAINT intervju_pk PRIMARY KEY ( idint );

CREATE TABLE korisnik (
    id      INTEGER NOT NULL,
    email   VARCHAR2(24 CHAR) NOT NULL,
    ime     VARCHAR2(24 CHAR) NOT NULL,
    prezime VARCHAR2(24 CHAR) NOT NULL,
    datrodj DATE NOT NULL,
    lozinka VARCHAR2(32 CHAR) NOT NULL,
    uloga   VARCHAR2(10) NOT NULL
);

ALTER TABLE korisnik
    ADD CHECK ( uloga IN ( 'Regrut', 'Zaposleni' ) );

ALTER TABLE korisnik ADD CONSTRAINT korisnik_pk PRIMARY KEY ( id );

CREATE TABLE kreira (
    idtim     INTEGER NOT NULL,
    idzeljtim INTEGER NOT NULL
);

ALTER TABLE kreira ADD CONSTRAINT kreira_pk PRIMARY KEY ( idtim,
                                                          idzeljtim );

CREATE TABLE menadzer (
    id INTEGER NOT NULL
);

ALTER TABLE menadzer ADD CONSTRAINT menadzer_pk PRIMARY KEY ( id );

CREATE TABLE nadgleda (
    idskaut INTEGER NOT NULL,
    idreg   INTEGER NOT NULL,
    idtrng  INTEGER NOT NULL
);

ALTER TABLE nadgleda
    ADD CONSTRAINT nadgleda_pk PRIMARY KEY ( idskaut,
                                             idreg,
                                             idtrng );

CREATE TABLE nedodirljivaimovinatima (
    idnedidtim      INTEGER NOT NULL,
    datdodnedimotim DATE NOT NULL,
    belesnedimotim  VARCHAR2(64 CHAR),
    idtipnedimo     INTEGER NOT NULL,
    idtim           INTEGER
);

CREATE UNIQUE INDEX nedodirljivaimovinatima__idx ON
    nedodirljivaimovinatima (
        idtim
    ASC );

ALTER TABLE nedodirljivaimovinatima ADD CONSTRAINT nedodirljivaimovinatima_pk PRIMARY KEY ( idnedidtim );

CREATE TABLE pik (
    idpik      INTEGER NOT NULL,
    redbrpik   VARCHAR2(2 CHAR) NOT NULL,
    brrunpik   VARCHAR2(1 CHAR) NOT NULL,
    godpik     VARCHAR2(4 CHAR) NOT NULL,
    idmenadzer INTEGER,
    idtim      INTEGER NOT NULL
);

ALTER TABLE pik ADD CONSTRAINT pik_pk PRIMARY KEY ( idpik );

CREATE TABLE pozivnaintervju (
    idpozint     INTEGER NOT NULL,
    mesodrpozint VARCHAR2(24 CHAR) NOT NULL,
    datvrepozint DATE NOT NULL,
    statuspozint VARCHAR2(10) NOT NULL,
    razodbpozint VARCHAR2(64 CHAR) NOT NULL,
    idregrut     INTEGER,
    idtrener     INTEGER NOT NULL
);

ALTER TABLE pozivnaintervju
    ADD CHECK ( statuspozint IN ( 'AFFIRMED', 'REJECTED', 'WAITING' ) );

ALTER TABLE pozivnaintervju ADD CONSTRAINT pozivnaintervju_pk PRIMARY KEY ( idpozint );

CREATE TABLE pozivnatrening (
    idpoztrng     INTEGER NOT NULL,
    datvrepoztrng DATE NOT NULL,
    mesodrpoztrng VARCHAR2(32 CHAR) NOT NULL,
    statuspoztrng VARCHAR2(10) NOT NULL,
    razodbpoztrng VARCHAR2(64 CHAR),
    idtrener      INTEGER NOT NULL
);

ALTER TABLE pozivnatrening
    ADD CHECK ( statuspoztrng IN ( 'AFFIRMED', 'REJECTED', 'WAITING' ) );

ALTER TABLE pozivnatrening ADD CONSTRAINT pozivnatrening_pk PRIMARY KEY ( idpoztrng );

CREATE TABLE pravanaigraca (
    idprava          INTEGER NOT NULL,
    imeigrprava      VARCHAR2(24 CHAR) NOT NULL,
    prezimeigrprava  VARCHAR2(24 CHAR) NOT NULL,
    pozicijaigrprava VARCHAR2(2),
    idtim            INTEGER NOT NULL,
    idregrut         INTEGER NOT NULL,
    idpik            INTEGER NOT NULL
);

ALTER TABLE pravanaigraca
    ADD CHECK ( pozicijaigrprava IN ( 'C', 'PF', 'PG', 'SF', 'SG' ) );

CREATE UNIQUE INDEX pravanaigraca__idx ON
    pravanaigraca (
        idregrut
    ASC,
        idpik
    ASC );

ALTER TABLE pravanaigraca ADD CONSTRAINT pravanaigraca_pk PRIMARY KEY ( idprava );

CREATE TABLE predmettrgovine (
    idpredtrg  INTEGER NOT NULL,
    tippredtrg VARCHAR2(20) NOT NULL,
    idprava    INTEGER,
    idigrac    INTEGER,
    idzahtrg   INTEGER NOT NULL,
    idpik      INTEGER
);

ALTER TABLE predmettrgovine
    ADD CHECK ( tippredtrg IN ( 'Igrac', 'Pik', 'PravaNaIgraca' ) );

ALTER TABLE predmettrgovine ADD CONSTRAINT predmettrgovine_pk PRIMARY KEY ( idpredtrg );

CREATE TABLE regrut (
    id          INTEGER NOT NULL,
    kontelefon  VARCHAR2(24 CHAR) NOT NULL,
    visreg      VARCHAR2(4 CHAR) NOT NULL,
    tezreg      VARCHAR2(6 CHAR) NOT NULL,
    pozreg      VARCHAR2(2) NOT NULL,
    prosrankreg VARCHAR2(6 CHAR) NOT NULL,
    prosocreg   VARCHAR2(6 CHAR) NOT NULL,
    iddraft     INTEGER
);

ALTER TABLE regrut
    ADD CHECK ( pozreg IN ( 'C', 'PF', 'PG', 'SF', 'SG' ) );

ALTER TABLE regrut ADD CONSTRAINT regrut_pk PRIMARY KEY ( id );

CREATE TABLE skaut (
    id INTEGER NOT NULL
);

ALTER TABLE skaut ADD CONSTRAINT skaut_pk PRIMARY KEY ( id );

CREATE TABLE tim (
    idtim     INTEGER NOT NULL,
    naztim    VARCHAR2(32 CHAR) NOT NULL,
    godosntim VARCHAR2(4 CHAR) NOT NULL,
    loktim    VARCHAR2(32 CHAR) NOT NULL
);

ALTER TABLE tim ADD CONSTRAINT tim_pk PRIMARY KEY ( idtim );

CREATE TABLE tipimovinezatrgovinu (
    idtipimotrg  INTEGER NOT NULL,
    naztipimotrg VARCHAR2(32 CHAR) NOT NULL
);

ALTER TABLE tipimovinezatrgovinu ADD CONSTRAINT tipimovinezatrgovinu_pk PRIMARY KEY ( idtipimotrg );

CREATE TABLE tipnedodirljiveimovine (
    idtipnedimo  INTEGER NOT NULL,
    naztipnedimo VARCHAR2(32 CHAR) NOT NULL
);

ALTER TABLE tipnedodirljiveimovine ADD CONSTRAINT tipnedodirljiveimovine_pk PRIMARY KEY ( idtipnedimo );

CREATE TABLE tiptreninga (
    idtiptrng   INTEGER NOT NULL,
    naztiptrng  VARCHAR2(32 CHAR) NOT NULL,
    ciljtiptrng VARCHAR2(64 CHAR) NOT NULL
);

ALTER TABLE tiptreninga ADD CONSTRAINT tiptreninga_pk PRIMARY KEY ( idtiptrng );

CREATE TABLE tipugovora (
    idtipugo  INTEGER NOT NULL,
    naztipugo VARCHAR2(24 CHAR) NOT NULL
);

ALTER TABLE tipugovora ADD CONSTRAINT tipugovora_pk PRIMARY KEY ( idtipugo );

CREATE TABLE tipzelje (
    idtipzelje  INTEGER NOT NULL,
    naztipzelje VARCHAR2(32 CHAR) NOT NULL
);

ALTER TABLE tipzelje ADD CONSTRAINT tipzelje_pk PRIMARY KEY ( idtipzelje );

CREATE TABLE trener (
    id           INTEGER NOT NULL,
    godisktrener VARCHAR2(2 CHAR) NOT NULL,
    spectrener   VARCHAR2(20) NOT NULL
);

ALTER TABLE trener
    ADD CHECK ( spectrener IN ( 'DEFENSE', 'OFFENSE', 'PLAYER_MANAGEMENT' ) );

ALTER TABLE trener ADD CONSTRAINT trener_pk PRIMARY KEY ( id );

CREATE TABLE trening (
    idtrng     INTEGER NOT NULL,
    trajtrng   VARCHAR2(4 CHAR) NOT NULL,
    datvretrng DATE NOT NULL,
    mesodrtrng VARCHAR2(32 CHAR) NOT NULL,
    belestrng  VARCHAR2(64 CHAR) NOT NULL,
    idtiptrng  INTEGER NOT NULL,
    idpoztrng  INTEGER NOT NULL
);

CREATE UNIQUE INDEX trening__idx ON
    trening (
        idpoztrng
    ASC );

ALTER TABLE trening ADD CONSTRAINT trening_pk PRIMARY KEY ( idtrng );

CREATE TABLE trgovina (
    idtrg    INTEGER NOT NULL,
    dattrg   DATE NOT NULL,
    tiptrg   VARCHAR2(20) NOT NULL,
    idzahtrg INTEGER NOT NULL
);

ALTER TABLE trgovina
    ADD CHECK ( tiptrg IN ( 'PICK_PICK', 'PLAYER_PICK', 'PLAYER_PLAYER' ) );

CREATE UNIQUE INDEX trgovina__idx ON
    trgovina (
        idzahtrg
    ASC );

ALTER TABLE trgovina ADD CONSTRAINT trgovina_pk PRIMARY KEY ( idtrg );

CREATE TABLE ucestvuje (
    idregrut INTEGER NOT NULL,
    idtrng   INTEGER NOT NULL
);

ALTER TABLE ucestvuje ADD CONSTRAINT ucestvuje_pk PRIMARY KEY ( idregrut,
                                                                idtrng );

CREATE TABLE ugovor (
    idugo     INTEGER NOT NULL,
    datpotugo DATE,
    datvazugo DATE NOT NULL,
    vredugo   VARCHAR2(6 CHAR) NOT NULL,
    opcugo    VARCHAR2(20) NOT NULL,
    idtim     INTEGER NOT NULL,
    idtipugo  INTEGER NOT NULL
);

ALTER TABLE ugovor
    ADD CHECK ( opcugo IN ( 'NO_OPTION', 'PLAYER_OPTION', 'TEAM_OPTION' ) );

ALTER TABLE ugovor ADD CONSTRAINT ugovor_pk PRIMARY KEY ( idugo );

CREATE TABLE upravlja (
    idregrut  INTEGER NOT NULL,
    idpoztrng INTEGER NOT NULL
);

ALTER TABLE upravlja ADD CONSTRAINT upravlja_pk PRIMARY KEY ( idregrut,
                                                              idpoztrng );

CREATE TABLE zahtevzatrgovinu (
    idzahtrg       INTEGER NOT NULL,
    datzahtrg      DATE NOT NULL,
    tipzahtrg      VARCHAR2(20) NOT NULL,
    statuszahtrg   VARCHAR2(20) NOT NULL,
    razlogodbij    VARCHAR2(64 CHAR),
    idmenadzerpos  INTEGER NOT NULL,
    idmenadzerprim INTEGER NOT NULL
);

ALTER TABLE zahtevzatrgovinu
    ADD CHECK ( tipzahtrg IN ( 'PICK_PICK', 'PLAYER_PICK', 'PLAYER_PLAYER' ) );

ALTER TABLE zahtevzatrgovinu
    ADD CHECK ( statuszahtrg IN ( 'ACCEPTED', 'CANCELLED', 'DECLINED', 'IN_PROGRESS' ) );

ALTER TABLE zahtevzatrgovinu ADD CONSTRAINT zahtevzatrgovinu_pk PRIMARY KEY ( idzahtrg );

CREATE TABLE zaposleni (
    id     INTEGER NOT NULL,
    ulozap VARCHAR2(10) NOT NULL,
    mbrzap VARCHAR2(32 CHAR) NOT NULL,
    idugo  INTEGER NOT NULL
);

ALTER TABLE zaposleni
    ADD CHECK ( ulozap IN ( 'Igrac', 'Menadzer', 'Skaut', 'Trener' ) );

CREATE UNIQUE INDEX zaposleni__idx ON
    zaposleni (
        idugo
    ASC );

ALTER TABLE zaposleni ADD CONSTRAINT zaposleni_pk PRIMARY KEY ( id );

CREATE TABLE zeljatima (
    idzeljtim     INTEGER NOT NULL,
    datdodzeljtim DATE NOT NULL,
    beleszeljtim  VARCHAR2(64 CHAR),
    idtipzelje    INTEGER NOT NULL
);

ALTER TABLE zeljatima ADD CONSTRAINT zeljatima_pk PRIMARY KEY ( idzeljtim );

ALTER TABLE bira
    ADD CONSTRAINT bira_pik_fk FOREIGN KEY ( idpik )
        REFERENCES pik ( idpik );

ALTER TABLE bira
    ADD CONSTRAINT bira_regrut_fk FOREIGN KEY ( idregrut )
        REFERENCES regrut ( id );

ALTER TABLE igrac
    ADD CONSTRAINT igrac_zaposleni_fk FOREIGN KEY ( id )
        REFERENCES zaposleni ( id );

ALTER TABLE imovinazatrgovinutima
    ADD CONSTRAINT imotrgtim_tim_fk FOREIGN KEY ( idtim )
        REFERENCES tim ( idtim );

ALTER TABLE imovinazatrgovinutima
    ADD CONSTRAINT imotrgtim_tipimotrg_fk FOREIGN KEY ( idtipimotrg )
        REFERENCES tipimovinezatrgovinu ( idtipimotrg );

ALTER TABLE intervju
    ADD CONSTRAINT intervju_pozivnaintervju_fk FOREIGN KEY ( idpozint )
        REFERENCES pozivnaintervju ( idpozint );

ALTER TABLE intervju
    ADD CONSTRAINT intervju_regrut_fk FOREIGN KEY ( idregrut )
        REFERENCES regrut ( id );

ALTER TABLE kreira
    ADD CONSTRAINT kreira_tim_fk FOREIGN KEY ( idtim )
        REFERENCES tim ( idtim );

ALTER TABLE kreira
    ADD CONSTRAINT kreira_zeljatima_fk FOREIGN KEY ( idzeljtim )
        REFERENCES zeljatima ( idzeljtim );

ALTER TABLE menadzer
    ADD CONSTRAINT menadzer_zaposleni_fk FOREIGN KEY ( id )
        REFERENCES zaposleni ( id );

ALTER TABLE nadgleda
    ADD CONSTRAINT nadgleda_skaut_fk FOREIGN KEY ( idskaut )
        REFERENCES skaut ( id );

ALTER TABLE nadgleda
    ADD CONSTRAINT nadgleda_ucestvuje_fk FOREIGN KEY ( idreg,
                                                       idtrng )
        REFERENCES ucestvuje ( idregrut,
                               idtrng );

ALTER TABLE nedodirljivaimovinatima
    ADD CONSTRAINT nedoimotim_tim_fk FOREIGN KEY ( idtim )
        REFERENCES tim ( idtim );

ALTER TABLE nedodirljivaimovinatima
    ADD CONSTRAINT nedoimotim_tipnedoimo_fk FOREIGN KEY ( idtipnedimo )
        REFERENCES tipnedodirljiveimovine ( idtipnedimo );

ALTER TABLE pik
    ADD CONSTRAINT pik_menadzer_fk FOREIGN KEY ( idmenadzer )
        REFERENCES menadzer ( id );

ALTER TABLE pik
    ADD CONSTRAINT pik_tim_fk FOREIGN KEY ( idtim )
        REFERENCES tim ( idtim );

ALTER TABLE pozivnaintervju
    ADD CONSTRAINT pozivnaintervju_regrut_fk FOREIGN KEY ( idregrut )
        REFERENCES regrut ( id );

ALTER TABLE pozivnaintervju
    ADD CONSTRAINT pozivnaintervju_trener_fk FOREIGN KEY ( idtrener )
        REFERENCES trener ( id );

ALTER TABLE pozivnatrening
    ADD CONSTRAINT pozivnatrening_trener_fk FOREIGN KEY ( idtrener )
        REFERENCES trener ( id );

ALTER TABLE pravanaigraca
    ADD CONSTRAINT pravanaigraca_bira_fk FOREIGN KEY ( idregrut,
                                                       idpik )
        REFERENCES bira ( idregrut,
                          idpik );

ALTER TABLE pravanaigraca
    ADD CONSTRAINT pravanaigraca_tim_fk FOREIGN KEY ( idtim )
        REFERENCES tim ( idtim );

ALTER TABLE predmettrgovine
    ADD CONSTRAINT predtrg_igrac_fk FOREIGN KEY ( idigrac )
        REFERENCES igrac ( id );

ALTER TABLE predmettrgovine
    ADD CONSTRAINT predtrg_pik_fk FOREIGN KEY ( idpik )
        REFERENCES pik ( idpik );

ALTER TABLE predmettrgovine
    ADD CONSTRAINT predtrg_pravanaigraca_fk FOREIGN KEY ( idprava )
        REFERENCES pravanaigraca ( idprava );

ALTER TABLE predmettrgovine
    ADD CONSTRAINT predtrg_zahzatrg_fk FOREIGN KEY ( idzahtrg )
        REFERENCES zahtevzatrgovinu ( idzahtrg );

ALTER TABLE regrut
    ADD CONSTRAINT regrut_draft_fk FOREIGN KEY ( iddraft )
        REFERENCES draft ( iddraft );

ALTER TABLE regrut
    ADD CONSTRAINT regrut_korisnik_fk FOREIGN KEY ( id )
        REFERENCES korisnik ( id );

ALTER TABLE skaut
    ADD CONSTRAINT skaut_zaposleni_fk FOREIGN KEY ( id )
        REFERENCES zaposleni ( id );

ALTER TABLE trener
    ADD CONSTRAINT trener_zaposleni_fk FOREIGN KEY ( id )
        REFERENCES zaposleni ( id );

ALTER TABLE trening
    ADD CONSTRAINT trening_pozivnatrening_fk FOREIGN KEY ( idpoztrng )
        REFERENCES pozivnatrening ( idpoztrng );

ALTER TABLE trening
    ADD CONSTRAINT trening_tiptreninga_fk FOREIGN KEY ( idtiptrng )
        REFERENCES tiptreninga ( idtiptrng );

ALTER TABLE trgovina
    ADD CONSTRAINT trgovina_zahtevzatrgovinu_fk FOREIGN KEY ( idzahtrg )
        REFERENCES zahtevzatrgovinu ( idzahtrg );

ALTER TABLE ucestvuje
    ADD CONSTRAINT ucestvuje_regrut_fk FOREIGN KEY ( idregrut )
        REFERENCES regrut ( id );

ALTER TABLE ucestvuje
    ADD CONSTRAINT ucestvuje_trening_fk FOREIGN KEY ( idtrng )
        REFERENCES trening ( idtrng );

ALTER TABLE ugovor
    ADD CONSTRAINT ugovor_tim_fk FOREIGN KEY ( idtim )
        REFERENCES tim ( idtim );

ALTER TABLE ugovor
    ADD CONSTRAINT ugovor_tipugovora_fk FOREIGN KEY ( idtipugo )
        REFERENCES tipugovora ( idtipugo );

ALTER TABLE upravlja
    ADD CONSTRAINT upravlja_pozivnatrening_fk FOREIGN KEY ( idpoztrng )
        REFERENCES pozivnatrening ( idpoztrng );

ALTER TABLE upravlja
    ADD CONSTRAINT upravlja_regrut_fk FOREIGN KEY ( idregrut )
        REFERENCES regrut ( id );

ALTER TABLE zahtevzatrgovinu
    ADD CONSTRAINT zahtevzatrgovinu_menadzer_fk FOREIGN KEY ( idmenadzerpos )
        REFERENCES menadzer ( id );

ALTER TABLE zahtevzatrgovinu
    ADD CONSTRAINT zahtevzatrgovinu_menadzer_fkv1 FOREIGN KEY ( idmenadzerprim )
        REFERENCES menadzer ( id );

ALTER TABLE zaposleni
    ADD CONSTRAINT zaposleni_korisnik_fk FOREIGN KEY ( id )
        REFERENCES korisnik ( id );

ALTER TABLE zaposleni
    ADD CONSTRAINT zaposleni_ugovor_fk FOREIGN KEY ( idugo )
        REFERENCES ugovor ( idugo );

ALTER TABLE zeljatima
    ADD CONSTRAINT zeljatima_tipzelje_fk FOREIGN KEY ( idtipzelje )
        REFERENCES tipzelje ( idtipzelje );

CREATE SEQUENCE SEQ_IDPREDTRG
INCREMENT BY 1
START WITH 1
NOCYCLE
CACHE 10;

CREATE OR REPLACE TRIGGER PREDMETTRGOVINE_PK_TRIGGER
BEFORE INSERT
ON PREDMETTRGOVINE
FOR EACH ROW
BEGIN
 SELECT Seq_IDPREDTRG.NEXTVAL
 INTO :NEW.IDPREDTRG
 FROM SYS.DUAL;
END;

CREATE SEQUENCE SEQ_IDZAHTRG
INCREMENT BY 1
START WITH 1
NOCYCLE
CACHE 10;

CREATE OR REPLACE TRIGGER ZAHTEVZATRGOVINU_PK_TRIGGER
BEFORE INSERT
ON ZAHTEVZATRGOVINU
FOR EACH ROW
BEGIN
 SELECT Seq_IDZAHTRG.NEXTVAL
 INTO :NEW.IDZAHTRG
 FROM SYS.DUAL;
END;

CREATE SEQUENCE SEQ_IDTRG
INCREMENT BY 1
START WITH 1
NOCYCLE
CACHE 10;

CREATE OR REPLACE TRIGGER TRGOVINA_PK_TRIGGER
BEFORE INSERT
ON TRGOVINA
FOR EACH ROW
BEGIN
 SELECT Seq_IDTRG.NEXTVAL
 INTO :NEW.IDTRG
 FROM SYS.DUAL;
END;

CREATE OR REPLACE TRIGGER ZAHTEVZATRGOVINU_ACCEPT_TR
FOR UPDATE OF STATUSZAHTRG
ON ZAHTEVZATRGOVINU
COMPOUND TRIGGER

  TYPE ZahtIDTableType IS TABLE OF ZAHTEVZATRGOVINU.IDZAHTRG%TYPE;
  ZahtIDs ZahtIDTableType := ZahtIDTableType();

  BEFORE EACH ROW IS
  BEGIN
    IF :NEW.STATUSZAHTRG = 'ACCEPTED' THEN
      ZahtIDs.EXTEND;
      ZahtIDs(ZahtIDs.COUNT) := :OLD.IDZAHTRG;
    END IF;
  END BEFORE EACH ROW;

  AFTER STATEMENT IS
  BEGIN
    FOR i IN 1 .. ZahtIDs.COUNT LOOP
      FOR PREDMET IN (SELECT P.IDPREDTRG, P.TIPPREDTRG, P.IDPRAVA, P.IDIGRAC, P.IDZAHTRG, P.IDPIK
                      FROM PREDMETTRGOVINE P
                      WHERE P.IDZAHTRG = ZahtIDs(i)) LOOP

        FOR ZAHTEV IN (SELECT Z.IDZAHTRG
                       FROM ZAHTEVZATRGOVINU Z
                       JOIN PREDMETTRGOVINE P ON P.IDZAHTRG = Z.IDZAHTRG
                       WHERE Z.STATUSZAHTRG = 'IN_PROGRESS'
                         AND (P.IDIGRAC = PREDMET.IDIGRAC OR 
                              P.IDPIK = PREDMET.IDPIK OR 
                              P.IDPRAVA = PREDMET.IDPRAVA)) LOOP
            
          UPDATE ZAHTEVZATRGOVINU 
          SET STATUSZAHTRG = 'DECLINED', RAZLOGODBIJ = 'Another trade proposal has been accepted'
          WHERE IDZAHTRG = ZAHTEV.IDZAHTRG;
        
        END LOOP;
      END LOOP;
    END LOOP;
  END AFTER STATEMENT;
END;
