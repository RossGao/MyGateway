/*
The file provide functionality to cache user login info in local memory.
Take use of groupcache module to implement the cache funciton.
1. When authentication finished, the cache will store the authorizaiton code.
2. When authoriztion finished, the cache will store all the user's login info.*/

package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/golang/groupcache"
)

// Method stores users authentication info.
func CacheAuthenticationInfo(authenInfo AuthenticateInfo) error{

}

// Method stores users authorization info.
func CacheAtuhorizationInfo(authorizeInfo AuthorizationInfo) error{

}

// Get authentication info by user id.
func GetCachedAuthenInfo(userId string) AuthenticateInfo error{

}

// Get authorization info by user id.
func GetCachedAuthorizeInfo(userId string) AuthorizationInfo error{

}
