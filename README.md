Gives the ability to construct web services just as config. 
manages state operations between different versions
provides an event based system that triggers actions from CRUD updates on resources.
Offer native yaml configured plugins, or allow injectable scripts or modules (go binaries?)

# api constructor

project that constructs api's given from simple inputs. 
dynamically creates backend and objects 
yaml config describing schema of data and http resources created. 
resource updated can trigger configured hooks between crud actions. an updated entity would notify subscriptions and trigger processes.

# Resource

has a set of properties 
has relationships 
has an http endpoint
can trigger events when updated

# Interfaces

Provide an API that can be plugged into

# Authentication

Need a role based model, where resources can be granted public or private. 
Endpoint required, client logs in and gets token. With token can be used to access endpoint. 

ACL: 
- public, read only operations on public resources
- private, read / write on permissioned resources
- admin, operations to the resource (i.e. updating resources, adding triggers, )


## libraries of interest

hashicorp memdb


## future 

client tools to develop web services (i.e. wyswyg)
client libraries (javascript most notably) to plugin to backend and generate methods and resources from endpoint.

## example usecase.

Terraform Script

- Deploy backend with terraform resources
- Deploy s3 bucket
- Upload site to bucket
- Front with cloudfront + dns 