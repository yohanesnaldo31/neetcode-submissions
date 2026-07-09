import (
	"slices"
)

type Node struct {
    Index int
    Name string
    Email string
    Neighbors[]*Node
}

func accountsMerge(accounts [][]string) [][]string {
   mapEmailToAccount := make(map[string]*Node)
   arrayOfNode := make([]*Node,0)

   for i:= 0; i<len(accounts);i++{
        arrayOfNode = append(arrayOfNode,&Node{
            Index: i,
            Name: accounts[i][0],
            Neighbors: make([]*Node,0),
        })
        for j := 1; j<len(accounts[i]);j++ {
            node, ok:= mapEmailToAccount[accounts[i][j]]
            if ok {
                node.Neighbors = append(node.Neighbors, arrayOfNode[i])
                arrayOfNode[i].Neighbors = append(arrayOfNode[i].Neighbors, node)
            } else {
                emailNode := &Node{
                    Email: accounts[i][j],
                    Neighbors: []*Node{arrayOfNode[i]},
                }
                mapEmailToAccount[accounts[i][j]] = emailNode
                arrayOfNode[i].Neighbors = append(arrayOfNode[i].Neighbors, emailNode)
            }
        }
   }  
    
    mapTraversedNode := map[string]struct{}{}
    var dfs func(node *Node, emailCombination *[]string)
    dfs = func(node *Node, emailCombination *[]string){
        _, ok := mapTraversedNode[node.Email+"#"+strconv.Itoa(node.Index)]
        if ok{
            return
        }

        if node.Email != "" {
            *emailCombination=append(*emailCombination, node.Email)
        }
        mapTraversedNode[node.Email+"#"+strconv.Itoa(node.Index)] = struct{}{}
        for _, nodeNeighbor := range node.Neighbors{
            dfs(nodeNeighbor, emailCombination)
        }
    }

    result := [][]string{}
    for i:=0;i<len(arrayOfNode);i++{
        _, ok := mapTraversedNode["#"+strconv.Itoa(i)]
        if ok{
            continue
        }

        emailCombination := make([]string,0)
        dfs(arrayOfNode[i],&emailCombination)
        
        slices.Sort(emailCombination)

        mergedAccounts := make([]string, 0) 
        mergedAccounts = append([]string{accounts[i][0]} , emailCombination...)
        
        
        result = append(result, mergedAccounts)
    }

    return result
}