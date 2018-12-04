package jobmgr


import (
    "errors"
    "log"
    "golang.org/x/crypto/ssh"
    "io/ioutil"
    "fmt"
)


type I_Machine interface {

    GetDomain() string
    GetPort() string
    GetUsername() string
    GetMode() string
    GetSSHKeyfile() string
    GetPassphrase() string
    GetPassword() string
    GetSudoPassword() string
    GetAuthConf() (*ssh.ClientConfig, string, error)
}



/*
    Define class: Machine implementing I_Machine
*/
type Machine struct {
    Domain string `json:"domain"`
    Port string `json:"port"`
    Username string `json:"username"`
    Mode string `json:"mode"`
    Credential Credential `json:"credential"`
    SudoPassword string `json:"sudopassword"`
}

type Credential struct {
    Passphrase string `json:"passphrase"`
    Password string `json:"password"`
    SSHKeyfile string `json:"sshkeyfile"`
}

/*
    Implement I_Machine.GetDomain()       Get machine domain name or ip 

*/
func (m *Machine) GetDomain() string {
    return m.Domain
}


/*
    Implement I_Machine.GetPort()       Get ssh port 

*/
func (m *Machine) GetPort() string {
    return m.Port
}

/*
    Implement I_Machine.GetUsername()       Get username

*/
func (m *Machine) GetUsername() string {
    return m.Username
}


/*
    Implement I_Machine.GetMode()       Get approach of logging into machine: SSHKEY, USERPASS

*/
func (m *Machine) GetMode() string {
    return m.Mode
}

/*
    Implement I_Machine.GetPassphrase()       Get passphrase 

*/
func (m *Machine) GetPassphrase() string {
    return m.Credential.Passphrase
}

/*
    Implement I_Machine.GetPassword()       Get password 

*/
func (m *Machine) GetPassword() string {
    return m.Credential.Password
}


/*
    Implement I_Machine.GetSSHKeyfile()       Get ssh private key file 

*/
func (m *Machine) GetSSHKeyfile() string {
    return m.Credential.SSHKeyfile
}


/*
    Implement I_Machine.GetSudoPassword()       Get sudo password of user

*/
func (m *Machine) GetSudoPassword() string {
    return m.SudoPassword
}

func (m *Machine) GetSigner() (ssh.Signer, error) {
    key, err := ioutil.ReadFile(m.GetSSHKeyfile())
    if err != nil {
        log.Fatalf("unable to read private key: %v", err)
    }
    // Create the Signer for this private key.
    switch mode := m.GetMode(); mode {
    case "SSHKEY":
        return ssh.ParsePrivateKey(key)
    case "SSHKEYWITHPASSPHRASE":
        return ssh.ParsePrivateKeyWithPassphrase(key, []byte(m.GetPassphrase()))
    default:
        return nil, errors.New("Mode is not supported")
    }
}

/*
    Implement I_Machine.GetAutoConf()       Get auth conf for ssh 

*/
func (m *Machine) GetAuthConf() (*ssh.ClientConfig, string, error) {
    var auth []ssh.AuthMethod
    var e error
    switch mode := m.GetMode(); mode {
    case "USERPASS":
        auth = []ssh.AuthMethod{
            ssh.Password(m.GetPassword()),
        }
        e = nil
    case "SSHKEY", "SSHKEYWITHPASSPHRASE":
        signer, err := m.GetSigner()
        if err != nil {
            log.Fatalf("unable to parse private key: %v", err)
        }
        auth = []ssh.AuthMethod{
            ssh.PublicKeys(signer),
        }
        e = err
    default:
        auth = []ssh.AuthMethod{}
        e = errors.New("Mode is not supported")
    }
    config := &ssh.ClientConfig{
        User: m.GetUsername(),
        Auth: auth,
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    dest := fmt.Sprintf("%s:%s", m.GetDomain(), m.GetPort())
    return config, dest, e
}
