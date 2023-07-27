# Copyright (c) 2023, Circle Technologies, LLC. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import os
import codecs

# The following sample codes generate a distinct hex encoded entity secret with each execution.
# The generation of entity secret only need to be executed once unless you need to rotate entity secret.
if __name__ == '__main__':
    entity_secret = os.urandom(32)
    print("Hex encoded entity secret:", codecs.encode(entity_secret, 'hex').decode())
