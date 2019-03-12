Getting Started
===============

.. toctree::
   :maxdepth: 1
   :hidden:

   prereqs
   install

Before we begin, if you haven't already done so, you may wish to check that
you have all the :doc:`prereqs` installed on the platform(s)
on which you'll be developing blockchain applications and/or operating
Hyperledger UDO.

Once you have the prerequisites installed, you are ready to download and
install HyperLedger UDO. While we work on developing real installers for the
UDO binaries, we provide a script that will :doc:`install` to your system.
The script also will download the Docker images to your local registry.


Hyperledger UDO SDKs
^^^^^^^^^^^^^^^^^^^^^^^

Hyperledger UDO offers a number of SDKs to support various programming
languages. There are two officially released SDKs for Node.js and Java:

  * `Hyperledger UDO Node SDK <https://github.com/hyperledger/udo-sdk-node>`__ and `Node SDK documentation <https://udo-sdk-node.github.io/>`__.
  * `Hyperledger UDO Java SDK <https://github.com/hyperledger/udo-sdk-java>`__.

In addition, there are three more SDKs that have not yet been officially released
(for Python, Go and REST), but they are still available for downloading and testing:

  * `Hyperledger UDO Python SDK <https://github.com/hyperledger/udo-sdk-py>`__.
  * `Hyperledger UDO Go SDK <https://github.com/hyperledger/udo-sdk-go>`__.
  * `Hyperledger UDO REST SDK <https://github.com/hyperledger/udo-sdk-rest>`__.

Hyperledger UDO CA
^^^^^^^^^^^^^^^^^^^^^

Hyperledger UDO provides an optional
`certificate authority service <http://hyperledger-fabric-ca.readthedocs.io/en/latest>`_
that you may choose to use to generate the certificates and key material
to configure and manage identity in your blockchain network. However, any CA
that can generate ECDSA certificates may be used.

.. Licensed under Creative Commons Attribution 4.0 International License
   https://creativecommons.org/licenses/by/4.0/
