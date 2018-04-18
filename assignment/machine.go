package assignment


import (
)


/*
    Define interface: I_Machine
        
        GetMode()       Get approach of logging into machine: SSHKEY, USERPASS
        GetLabel()      Get Label of the machine
        GetSudoPassword()       Get sudo password of user
*/
type I_Machine interface {
    GetMode() string
    GetLabel() string
    GetSudoPassword() string
    GetDomain() string
    GetUsername() string
    GetPort() string
}



/*
    Define class: Machine implementing I_Machine
*/
type Machine struct {
    Label string
    Domain string
    Username string
    Port string
    SudoPassword string
    Mode string
    CertPassword string  // NOT USED
}


/*
    Implement I_Machine.GetLabel()      Get Label of the machine

*/
func (m *Machine) GetLabel() string {
    return m.Label
}

/*
    Implement I_Machine.GetMode()       Get approach of logging into machine: SSHKEY, USERPASS

*/
func (m *Machine) GetMode() string {
    return m.Mode
}

/*
    Implement I_Machine.GetSudoPassword()       Get sudo password of user

*/
func (m *Machine) GetSudoPassword() string {
    return m.SudoPassword
}

/*
    Implement I_Machine.GetUsername()       Get username

*/
func (m *Machine) GetUsername() string {
    return m.Username
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

