/**
 * Copyright 2018 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package controller

import (
	"fmt"
	"os"
)

// TODO need to remove this error, since its moved to ubiquity it self
type NoMounterForVolumeError struct {
	mounter string
}

func (e *NoMounterForVolumeError) Error() string {
	return fmt.Sprintf("Mounter not found for backend: %s", e.mounter)
}

const MissingWwnMountRequestErrorStr = "volume related to scbe backend must have mountRequest.Opts[Wwn] not found (expect to have wwn for scbe backend volume type)"

const FailRemovePVorigDirErrorStr = "Failed removing existing volume directory"

type FailRemovePVorigDirError struct {
	err string
	dir error
}

func (e *FailRemovePVorigDirError) Error() string {
	return fmt.Sprintf(FailRemovePVorigDirErrorStr+". dir=[%s], err=[%#v]", e.dir, e.err)
}

const WrongSlinkErrorStr = "Idempotent - The existing slink point to a wrong mountpoint."

type wrongSlinkError struct {
	slink           string
	wrongPointTo    string
	expectedPointTo string
}

func (e *wrongSlinkError) Error() string {
	return fmt.Sprintf(WrongSlinkErrorStr+" slink=[%s] expected-mountpoint=[%s], wrong-mountpoint=[%s]", e.slink, e.expectedPointTo, e.wrongPointTo)
}

const SupportK8sVesion = "Support only k8s version >= 1.6."

type k8sVersionNotSupported struct {
	version string
}

func (e *k8sVersionNotSupported) Error() string {
	return fmt.Sprintf(SupportK8sVesion+" Version [%s] is not supported.", e.version)
}

const K8sPVDirectoryIsNotDirNorSlinkErrorStr = "k8s PV directory, k8s-mountpoint, is not a directory nor slink."

type k8sPVDirectoryIsNotDirNorSlinkError struct {
	slink    string
	fileInfo os.FileInfo
}

func (e *k8sPVDirectoryIsNotDirNorSlinkError) Error() string {
	// The error string contains also the fileInfo for debug purpose, in order to identify for example what is the actual FileInfo.Mode().
	return fmt.Sprintf(K8sPVDirectoryIsNotDirNorSlinkErrorStr+" slink=[%s], fileInfo=[%#v]",
		e.slink, e.fileInfo)
}

const IdempotentUnmountSkipOnVolumeNotExistWarnigMsg = "Unmount operation requested to work on not exist volume. Assume its idempotent issue - so skip Unmount."

const K8sPVDirectoryIsNotSlinkErrorStr = "k8s PV directory, k8s-mountpoint, is not slink."

type k8sPVDirectoryIsNotSlinkError struct {
	slink    string
	fileInfo os.FileInfo
}

func (e *k8sPVDirectoryIsNotSlinkError) Error() string {
	// The error string contains also the fileInfo for debug purpose, in order to identify for example what is the actual FileInfo.Mode().
	return fmt.Sprintf(K8sPVDirectoryIsNotSlinkErrorStr+" slink=[%s], fileInfo=[%#v]",
		e.slink, e.fileInfo)
}

const PvBackendNotSupportedErrorStr = "Backend type not supported."

type PvBackendNotSupportedError struct {
	Backend string
}

func (e *PvBackendNotSupportedError) Error() string {
	return fmt.Sprintf(PvBackendNotSupportedErrorStr+" backend=[%s]", e.Backend)
}

const BackendNotImplementedGetRealMountpointErrorStr = "Backend do not support getting the real mountpoint from PV"

type BackendNotImplementedGetRealMountpointError struct {
	Backend string
}

func (e *BackendNotImplementedGetRealMountpointError) Error() string {
	return fmt.Sprintf(BackendNotImplementedGetRealMountpointErrorStr+" backend=[%s]", e.Backend)
}

const PVIsAlreadyUsedByAnotherPodMessage = "PV is already in use by another pod and has an existing slink to mountpoint."

type PVIsAlreadyUsedByAnotherPod struct {
	mountpoint string
	slink      []string
}

func (e *PVIsAlreadyUsedByAnotherPod) Error() string {
	return fmt.Sprintf(PVIsAlreadyUsedByAnotherPodMessage + " mountpoint=[%s], slinks=[%s]", e.mountpoint, e.slink)
}

var WrongK8sDirectoryPathErrorMessage = fmt.Sprintf("Expected to find [%s] directory in k8s mount path.",K8sPodsDirecotryName)

type WrongK8sDirectoryPathError struct {
	k8smountdir string
}

func (e *WrongK8sDirectoryPathError) Error() string {
	return fmt.Sprintf(PVIsAlreadyUsedByAnotherPodMessage + "k8smountdir=[%s]", e.k8smountdir)
}

