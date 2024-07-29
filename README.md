# Ecommerce-product-listing

## System Architecture:

                              __________
                              (        )
                              (Admin DB)
                              (        )
                              ----------
                                  |
                                  |                ___________________
                                  |                (                 )
User <----> [Frontend] <----> [Backend] <--------> (Catalog Db Master)
                                  |     \          (                 )
                                  |       \        -------------------
                                  |         \           (       )
                              [Admin_UI]      \       (replication)
                                                \       (       )
                                                   ____________________
                                                   (                  )
                                                   (Catalog Db Replica)
                                                   (                  )     
                                                   --------------------



![system architecture](extras/system_architecture.png?raw=true)