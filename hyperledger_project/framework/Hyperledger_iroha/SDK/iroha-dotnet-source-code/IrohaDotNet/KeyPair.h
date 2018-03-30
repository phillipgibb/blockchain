/**
  * Copyright Jin Adachi All Rights Reserved.
  * https://github.com/hyperledger/iroha-dotnet
  * Licensed under the Apache License, Version 2.0 (the "License");
  * you may not use this file except in compliance with the License.
  * You may obtain a copy of the License at
  *      http://www.apache.org/licenses/LICENSE-2.0
  * Unless required by applicable law or agreed to in writing, software
  * distributed under the License is distributed on an "AS IS" BASIS,
  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  * See the License for the specific language governing permissions and
  * limitations under the License.
  */
#pragma once

using namespace System;

namespace IrohaDotNet
{
	public ref class KeyPair sealed
	{
	public:
		KeyPair(String ^privateKey, String ^publicKey)
		{
			PrivateKey = privateKey;
			PublicKey = publicKey;
		}
		property String ^PrivateKey;
		property String ^PublicKey;
	};
}